package test_test

import (
	"os/exec"

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
			cmd := exec.Command("bash", "-c", "eval", "$(crc podman-env)")
			_, err := cmd.Output()
			Expect(err).NotTo(HaveOccurred())
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
