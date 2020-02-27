package main

import (
	"golang.org/x/sys/windows/svc"
	//"golang.org/x/sys/windows/svc/debug"
	"io/ioutil"
	"time"
)

type dummyService struct {
}

func (srv *dummyService) Execute(args []string, req <-chan svc.ChangeRequest, stat chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	stat <- svc.Status{State: svc.StartPending}

	stopChan := make(chan bool, 1)
	go runBody(stopChan)

	stat <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

LOOP:
	for {
		switch r := <-req; r.Cmd {
		case svc.Stop, svc.Shutdown:
			stopChan <- true
			break LOOP

		case svc.Interrogate:
			stat <- r.CurrentStatus
			time.Sleep(100 * time.Millisecond)
			stat <- r.CurrentStatus
			//case svc.Pause:
			//case svc.Continue:
		}
	}

	stat <- svc.Status{State: svc.StopPending}
	return
}

func runBody(stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			return
		default:
			time.Sleep(10 * time.Second)
			ioutil.WriteFile("/temp/log.txt", []byte(time.Now().String()), 0)
		}
	}
}

//Windows Service Example
func main() {
	err := svc.Run("DummyService", &dummyService{})
	//err := debug.Run("DummyService", &dummyService{})
	if err != nil {
		panic(err)
	}
}

// 1.go get golang.org/x/sys/windows/svc
// 2.go build dummyService.go
// 3.sc create DummyService binPath= c:\goapp\src\dummyService.exe
// 4.sc delete DummyService
