package systray

import (
	"net"

	"github.com/Jacobious52/AnotherGamepad/http"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	log "github.com/sirupsen/logrus"
)

func LaunchTray(s *http.Server) {
	onExit := func() {
		log.Infoln("Cleaning up")
		s.Close()
	}
	systray.Run(onReady, onExit)
}

func openQR(i string) {

}

func getIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "[ERROR] Could not locate IP!"
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Another Gamepad")
	systray.SetTooltip("the cool tooltip")
	ip := getIp()

	mIp := systray.AddMenuItem(ip, "Code to type into your mobile device")
	mIp.SetIcon(icon.Data)
	go func() {
		<-mIp.ClickedCh
		openQR(ip)
	}()

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit Another Gamepad")
	mQuit.SetIcon(icon.Data)
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}
