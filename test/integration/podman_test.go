package test_test

import (
	"fmt"
	"os"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("podman preset", func() {

	Describe("basic use", func() {

		It("write to config", func() {
			Expect(RunCRCExpectSuccess("config", "set", "preset", "podman")).To(ContainSubstring("Changes to configuration property 'preset' are only applied when the CRC instance is created."))
		})

		It("setup CRC", func() {
			if bundlePath == "" {
				Expect(RunCRCExpectSuccess("setup")).To(ContainSubstring("Your system is correctly setup for using CodeReady Containers"))
			} else {
				Expect(RunCRCExpectSuccess("setup", "-b", bundlePath)).To(ContainSubstring("Your system is correctly setup for using CodeReady Containers"))
			}
		})

		It("start CRC", func() {
			if bundlePath == "" {
				Expect(RunCRCExpectSuccess("start")).To(ContainSubstring("podman runtime is now running"))
			} else {
				Expect(RunCRCExpectSuccess("start", "-b", bundlePath)).To(ContainSubstring("podman runtime is now running"))
			}
		})

		It("podman-env", func() {
			// Do what 'eval $(crc podman-env) would do
			path := os.ExpandEnv("${HOME}/.crc/bin/oc:$PATH")
			csshk := os.ExpandEnv("${HOME}/.crc/machines/crc/id_ecdsa")
			dh := os.ExpandEnv("unix:///${HOME}/.crc/machines/crc/docker.sock")
			ch := "ssh://core@127.0.0.1:2222/run/user/1000/podman/podman.sock"
			if runtime.GOOS == "windows" {
				dh = "npipe:////./pipe/rc-podman"
			}
			if runtime.GOOS == "linux" {
				ch = "ssh://core@192.168.130.11:22/run/user/1000/podman/podman.sock"
			}

			fmt.Println(path)
			fmt.Println(csshk)
			fmt.Println(ch)
			fmt.Println(dh)

			os.Setenv("PATH", path)
			os.Setenv("CONTAINER_SSHKEY", csshk)
			os.Setenv("CONTAINER_HOST", ch)
			os.Setenv("DOCKER_HOST", dh)
		})

		It("version", func() {
			Expect(RunPodmanExpectSuccess("version")).Should(MatchRegexp(`Version:[\s]*3\.\d+\.\d+`))
		})

		It("pull image", func() {
			RunPodmanExpectSuccess("pull", "fedora")
		})

		It("run image", func() {
			RunPodmanExpectSuccess("run", "fedora")
		})

		It("stop CRC", func() {
			Expect(RunCRCExpectSuccess("stop", "-f")).To(MatchRegexp("[Ss]topped the instance"))
		})
	})
})
