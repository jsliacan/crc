@draft_proxy
Feature: 
    End-to-end test to verify CRC use behind a proxy server.

    Scenario Outline: Set up CRC for use behind proxy
        When setting CRC config parameter "<parameter>" to "<value>" succeeds
        And that CRC setup finishes successfully

        Examples:
        | parameter   | value                        |
        | http-proxy  | http://squid.redhat.com:3128 |
        | https-proxy | http://squid.redhat.com:3128 |
        | no-proxy    | .testing                     |

    Scenario: Create project, create app
        Given that "config view" lists "squid.redhat.com:3128" as a proxy
        When CRC starts successfully and all requests during the start from "192.168.130.11" go via "squid.redhat.com" on port "3128"
        And running "oc-env" command succeeds
        Then with up to a few retries with wait period of a couple of mins, CRC status should be "Running"
        And login to the OpenShift cluster is successful
        When creating new-app via "oc" succeeds and all requests went through "squid.redhat.com" on port "3128"
        Then proxy works!

    Scenario: Clean-up
        Given deleting project succeeds and all requests from "192.168.130.11" go via "squid.redhat.com" on port "3128"
        When stopping and deleting CRC succeeds
        And running "crc cleanup" command succeeds
        Then life is good
