@performance @darwin @linux @windows

Feature: CRC performance status

    @cleanup @stability
    Scenario: Repeatedly start the cluster and observe if it is stable within 10mins after start
        Given executing single crc setup command succeeds
        Then run start-delete on repeat "50" times with "10m" cluster availability requirement

    @cleanup @running-state
    Scenario: Repeatedly start the cluster and observe if the state is 'Running' right away
        Given executing single crc setup command succeeds
        Then run start-delete on repeat "100" times checking for immediate running state
