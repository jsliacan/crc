@story_registry @darwin @linux @windows
Feature: Local image to image-registry to deployment

    The user creates a local container image with an app. They then
    push it to the openshift image-registry in their
    project/namespace. They deploy and expose the app and check its
    accessibility.

    Scenario: Start CRC
        Given executing "crc setup" succeeds
        When starting CRC with default bundle succeeds
        Then stdout should contain "Started the OpenShift cluster"
        And executing "eval $(crc oc-env)" succeeds
        When with up to "4" retries with wait period of "2m" command "crc status --log-level debug" output matches ".*Running \(v\d+\.\d+\.\d+.*\).*"
        Then login to the oc cluster succeeds

    Scenario: Create local image
        Given executing "cd ../../../testdata" succeeds
        When executing "sudo docker build -t hello:test ." succeeds
        And executing "cd ../integration" succeeds
        Then executing "sudo docker images" succeeds
        And stdout should contain "localhost/hello"
        
    Scenario: Push local image to OpenShift image registry
        Given executing "oc new-project testproj-img" succeeds
        When executing "sudo docker login -u kubeadmin -p $(oc whoami -t) default-route-openshift-image-registry.apps-crc.testing --tls-verify=false" succeeds
        Then stdout should contain "Login Succeeded!"
        When executing "sudo docker push hello:test default-route-openshift-image-registry.apps-crc.testing/testproj-img/hello:test --tls-verify=false" succeeds

    Scenario: Deploy the image
        Given executing "oc new-app testproj-img/hello:test" succeeds
        When executing "oc rollout status dc/hello" succeeds
        Then stdout should contain "successfully rolled out"
        When executing "oc get pods" succeeds
        Then stdout should contain "Running"
        And stdout should contain "Completed"
        When executing "oc logs -f dc/hello" succeeds
        Then stdout should contain "Hello, it works!"

    Scenario: Clean up
        Given executing "sudo docker images" succeeds
        When stdout contains "localhost/hello"
        Then executing "sudo docker image rm localhost/hello:test" succeeds
        And executing "oc delete project testproj-img" succeeds
        When executing "crc stop -f" succeeds
        Then stdout should match "(.*)[Ss]topped the OpenShift cluster"
        And executing "crc delete -f" succeeds
        Then stdout should contain "Deleted the OpenShift cluster"

