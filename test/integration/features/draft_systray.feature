@draft_systray
Feature: 

    Content of manual tests to verify system tray functionality on
    Windows and Mac OS.

    @darwin @windows
    Scenario Outline: Config
        When executing "crc config set" with "pull-secret-file" and "<pull-secret-file location>"
        And executing "crc config set" with "bundle" and "<bundle location>"
        Then executing "crc config view" should yield
        """
        - bundle                                : /path/to/bundle/crc_hypervisor_version.crcbundle
        - pull-secret-file                      : /path/to/pullsecret//crc-pull-secret
        """

    @darwin @windows
    Scenario Outline: Set-up
        When executing "crc setup --enable-experimental-features" goes well
        Then output should resemble
        """
        INFO Checking if oc binary is cached              
        INFO Caching oc binary                            
        ...
        Setup is complete, you can now run 'crc start -b $bundlename' to start the OpenShift cluster
        """

    @darwin @windows
    Scenario: Use tray to control CRC
        Given that tray icon is present in the system tray
        When pressing "Start" in the tray
        Then tray shows status "Starting" and within "10m" starts CRC
        And CRC status in the tray and in CLI is "Running" and detailed status is populated with correct information
        When pressing "About" section of the tray
        Then correct About info is displayed
        When pressing "Stop" in the tray
        Then CRC status in the tray and in CLI is "Stopped" in due time
        When pressing "Start" in the tray
        Then CRC status in the tray and in CLI is "Running" in due time
        When executing "crc stop" goes well in CLI
        Then CRC status in the tray and in CLI is "Stopped" in due time

    @darwin @windows
    Scenario: Clean-up
        When stopping and deleting CRC succeeds
        And running "crc cleanup" command succeeds
        Then life is good
