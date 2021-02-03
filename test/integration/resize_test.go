package test_test

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"runtime"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const testTopic = "resize"

type testTopic struct {
	Description string `json:"description"`
	Cases       []Case `json:"cases"`
}

type Case struct {
	Name        string     `json:"name"`
	Success     bool       `json:"success"`
	ErrorString string     `json:"error"`
	Stdout      string     `json:"stdout"`
	Parameters  Parameters `json:"parameters"`
	Checks      Validation `json:"checks"`
}

type Validation struct {
	Expected  Parameters `json:"expected"`
	Tolerance Parameters `json:"tolerance"`
}

type Parameters struct {
	CPU    string `json:"cpu"`
	Disk   string `json:"disk"`
	Memory string `json:"memory"`
}

var _ = Describe("changing VM properties: cpus, disk, memory", func() {

	byteInput, _ := ioutil.ReadAll(testTopic + "_test.json")
	var resizeTopic testTopic // holds test declaration
	json.Unmarshal(resizeTopic, byteInput)

	Describe(resizeTopic.Description, func() {

		It("setup CRC", func() {
			Expect(RunCRCExpectSuccess("setup")).To(ContainSubstring("Setup is complete"))
		})

		testCase := resizeTopic.Cases[0]
		// start + stop CRC for each case
		It("start CRC with "+testCase.Name, func() {
			if testCase.Success {
				Expect(RunCRCExpectSuccess("start", "-b", bundleLocation, "-p", pullSecretLocation, "--memory", testCase.Parameters.Memory, "--cpus", testCase.Parameters.CPU, "--disk-size", testCase.Parameters.Disk)).To(ContainSubstring(testCase.Stdout))

				It("check if memory was allocated correctly", func() {
					out, err := SendCommandToVM("lsmem --summary -b")
					Expect(err).NotTo(HaveOccurred())
					re := regexp.MustCompile("Total online memory: *[0-9]+")
					memoryLine := re.FindAllString(out, 1)[0]
					memoryBytes := strings.Atoi(strings.Fields(memoryLine)[3])

					expectedMem := strconv.Atoi(testCase.Checks.Expected.Memory)
					toleranceMem := strconv.Atoi(testCase.Checks.Tolerance.Memory)
					Expect(memoryBytes).Should(BeNumerically("<", expectedMem+toleranceMem))
					Expect(memoryBytes).Should(BeNumerically(">", expectedMem-toleranceMem))
				})

				It("check if CPU was allocated correctly", func() {
					out, err := SendCommandToVM("cat /proc/cpuinfo")
					Expect(err).NotTo(HaveOccurred())
					re := regexp.MustCompile("processor *: *[0-9]+")
					cpuLine := re.FindAllString(out, 1)[testCase.Checks.Expected.CPU]
					memoryBytes := strings.Atoi(strings.Fields(memoryLine)[3])
					Expect(out).Should(MatchRegexp(``))
				})

				// only check disk size on linux and windows
				if os := runtime.GOOS; os == "linux" || os == "windows" {

					It("check size of VM disk", func() {
						out, err := SendCommandToVM("df -h")
						Expect(err).NotTo(HaveOccurred())
						Expect(out).Should(MatchRegexp(`.*coreos-luks-root-nocrypt[\s]*40G`))
					})
				} else { // darwin
					It("check size of VM disk", func() {
						out, err := SendCommandToVM("df -h")
						Expect(err).NotTo(HaveOccurred())
						Expect(out).Should(MatchRegexp(`.*coreos-luks-root-nocrypt[\s]*31G`)) // default
					})

				}

			} else {
				Expect(RunCRCExpectFail("start", "-b", bundleLocation, "-p", pullSecretLocation, "--memory", c.Parameters.Memory, "--cpus", c.Parameters.CPU, "--disk-size", c.Parameters.Disk)).To(ContainSubstring(c.ErrorString))
			}
		})

		// check if start applied the parameters

		//---------------------------------------------------------- FORMER CODE -------------------------------------------------------

		It("stop CRC", func() {
			Expect(RunCRCExpectSuccess("stop", "-f")).To(ContainSubstring("Stopped the OpenShift cluster"))
		})

		// try bad things
		It("start CRC with too little memory", func() { // less than min = 9216
			Expect(RunCRCExpectFail("start", "--memory", "9000")).To(ContainSubstring("requires memory in MiB >= 9216"))
		})
		It("start CRC with too few cpus", func() { // fewer than min
			Expect(RunCRCExpectFail("start", "--cpus", "3")).To(ContainSubstring("")) // TODO
		})
		It("start CRC and shrink disk", func() { // bigger than default && smaller than current
			Expect(RunCRCExpectFail("start", "--disk-size", "35")).To(ContainSubstring("")) // TODO: diff between darwin & the rest
		})
		It("start CRC and shrink disk", func() { // smaller than min = default = 31GiB
			Expect(RunCRCExpectFail("start", "--disk-size", "30")).To(ContainSubstring("")) // TODO: diff between darwin & the rest
		})

		// start with default specs
		It("start CRC with memory size and cpu count", func() {
			Expect(RunCRCExpectSuccess("start", "-b", bundleLocation, "--memory", "9216", "--cpus", "4")).To(ContainSubstring("Started the OpenShift cluster"))
		})

		It("check memory size", func() {
			out, err := SendCommandToVM("cat /proc/meminfo")
			Expect(err).NotTo(HaveOccurred())
			Expect(out).Should(MatchRegexp(`MemTotal:[\s]*8\d{6}`))
		})

		It("check number of cpus", func() {
			out, err := SendCommandToVM("cat /proc/cpuinfo")
			Expect(err).NotTo(HaveOccurred())
			Expect(out).Should(MatchRegexp(`processor[\s]*\:[\s]*3`))
			Expect(out).ShouldNot(MatchRegexp(`processor[\s]*\:[\s]*4`))
		})

		// only check disk size on linux and windows
		if os := runtime.GOOS; os == "linux" || os == "windows" {
			It("check size of VM disk", func() {
				out, err := SendCommandToVM("df -h")
				Expect(err).NotTo(HaveOccurred())
				Expect(out).Should(MatchRegexp(`.*coreos-luks-root-nocrypt[\s]*40G`)) // no change
			})
		}

		It("clean up", func() {
			RunCRCExpectSuccess("stop", "-f")
			RunCRCExpectSuccess("delete", "-f")
			RunCRCExpectSuccess("cleanup")

		})
	})
})
