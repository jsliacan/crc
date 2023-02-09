@performance @darwin @linux @windows

Feature: CRC performance status

    @cleanup
    Scenario:
        Given executing single crc setup command succeeds
        Then run start-delete on repeat "20" times with "10m" cluster availability requirement