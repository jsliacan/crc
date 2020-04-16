@draft_cert-recovery
Feature: 
    End-to-end test to verify that expired certificates renew automatically

    Scenario: Fake future and start CRC with old certificates
        Given no CRC instance exists
        When Network Time Protocol (NTP) and port (UDP/123) are blocked
        When changing date to more than "1 month" in the future succeeds
        Then CRC setup succeeds
        And CRC start succeeds
        And stdout will contain "Certificates renewed successfully"

    Scenario: Clean up
        Given that CRC stops successfully
        When CRC delete succeeds
        And CRC cleanup succeeds
        Then open UDP/123 and check that date and time are correct

    Scenario: Start CRC with old certificates, then switch to future
        Given no CRC instance exists
        When CRC setup succeeds
        Then CRC start succeeds
        When Network Time Protocol (NTP) and port (UDP/123) are blocked
        And changing date to more than "1 month" in the future succeeds
        Then stdout will contain "Certificates renewed successfully"

    Scenario: Clean up
        Given that CRC stops successfully
        When CRC delete succeeds
        And CRC cleanup succeeds
        Then open UDP/123 and check that date and time are correct
