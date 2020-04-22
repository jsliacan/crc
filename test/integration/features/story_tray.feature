@story_tray
Feature: 

    Check if system tray works on Mac. Only basic checks as it is just
    a preview.

    @darwin
    Scenario: Start CRC
        Given executing "crc setup" succeeds
        When starting CRC with default bundle succeeds
        Then stdout should contain "Started the OpenShift cluster"
        And executing "eval $(crc oc-env)" succeeds
        When with up to "4" retries with wait period of "2m" command "crc status --log-level debug" output matches ".*Running \(v\d+\.\d+\.\d+.*\).*"
        Then login to the oc cluster succeeds

    @windows
    Scenario: Start CRC on Windows
        Given executing "crc setup" succeeds
        When starting CRC with default bundle and nameserver "10.75.5.25" succeeds
        Then stdout should contain "Started the OpenShift cluster"
        And executing "crc oc-env | Invoke-Expression" succeeds
        When with up to "4" retries with wait period of "2m" command "crc status --log-level debug" output matches ".*Running \(v\d+\.\d+\.\d+.*\).*"
        Then login to the oc cluster succeeds

    @linux @darwin @windows    
    Scenario: Check cluster health
        Given executing "crc status" succeeds
        And stdout should match ".*Running \(v\d+\.\d+\.\d+.*\).*"
        When executing "oc get nodes"
        Then stdout contains "Ready" 
        And stdout does not contain "Not ready"
        # next line checks similar things as `crc status` except gives more informative output
        And with up to "5" retries with wait period of "1m" all cluster operators are running
