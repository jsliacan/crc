package test_test

import (
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Counting slow starts", Label("openshift-preset", "slow-start-count"), func() {

	Describe("", func() {

		It("setup", func() {
			if bundlePath == "" {
				Expect(RunCRCExpectSuccess("setup")).To(ContainSubstring("Your system is correctly setup for using CRC"))
			} else {
				Expect(RunCRCExpectSuccess("setup", "-b", bundlePath)).To(ContainSubstring("Your system is correctly setup for using CRC"))
			}
		})

		It("start and status", func() {

			countOther := 0
			countSlow := 0

			f, err := os.Create("/tmp/slow-start.dat")
			if err != nil {
				fmt.Println(err)
			}

			// repeat starts
			for i := 0; i < 100; i++ {
				if bundlePath == "" {
					RunCRCExpectSuccess("start", "-p", pullSecretPath)
				} else {
					RunCRCExpectSuccess("start", "-b", bundlePath, "-p", pullSecretPath)
				}

				// keep checking for up to 10min after start whether status == "Running"
				stable := 0
				for j := 0; j < 13; j++ {

					stdout := RunCRCExpectSuccess("status")

					if strings.Contains(stdout, "Starting") {
						stable = 0
						time.Sleep(60 * time.Second)
						fmt.Println("Still starting...")
						if j >= 10 {
							break
						}
					} else {
						stable += 1
					}

					if stable == 3 {
						countOther += 1
						break
					}

				}
				if stable < 3 {
					countSlow += 1
				}

				_ = RunCRCExpectSuccess("delete", "-f")
				s := fmt.Sprintf("countSlow: %d, countOther: %d\n", countSlow, countOther)
				f.WriteString(s)
				// also print into stdout
				fmt.Printf("countSlow: %d, countOther: %d\n", countSlow, countOther)
			}
		})

		It("cleanup", func() {
			Expect(RunCRCExpectSuccess("cleanup")).To(ContainSubstring("Cleanup finished"))
		})

	})
})
