@linux @dummy
Feature:
    Dummy

    Scenario: Bla1
        When executing "ls" succeeds
        Then stdout should contain "."

    Scenario: Bla2
        When executing "ls -al" succeeds
        Then stdout should not contain "lskjfldskjflsjfslfkj"

