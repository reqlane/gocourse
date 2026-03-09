package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// === SIGINT
// > Ctrl + C
// === SIGTERM
// > taskkill //pid 123456
// === SIGKILL
// > taskkill //f //pid 123456

func main() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		done <- struct{}{}
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal")
				os.Exit(0)
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}
		// sig := <-sigs
		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Received SIGINT (Interrupt)")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Received SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Received SIGHUP (Hangup)")
		// 		// case syscall.SIGUSR1:
		// 		// 	fmt.Println("Received SIGUSR1 (User defined Signal 1)")
		// 		// 	fmt.Println("User defined function is executed")
		// 		// 	continue
		// 	}
		// 	fmt.Println("Graceful exit")
		// 	os.Exit(0)
		// }
	}()

	// Simulate work
	// fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}
