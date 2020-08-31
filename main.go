package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		command = "start"
	}

	switch command {
	case "start":
		start()
	case "child":
		setupNS()
		child()
	default:
		panic("oh oh")
	}
}

func start() {
	cmd := exec.Command("/proc/self/exe", "child") // symlink to currently running executable

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,           // 0 is root
				HostID:      os.Getuid(), // root user of the new user ns is mapped to the host's user ns
				Size:        1,           // range of IDs to map
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	cmd.Run()
}

func child() {
	// syscall.Chroot("/home/pravin")
	// syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command("/bin/bash")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

}

func setupNS() {
	// the new ns has been created at this point
	// allowing the new process to run in the correctly set hostname
	syscall.Sethostname([]byte("container"))
}
