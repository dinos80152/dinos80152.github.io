# Continuous Delivery

## Sequence

```sequence
local->git repo: git push
git repo->integration: request
integration->git repo: git pull
git repo-->integration: source code
integration->integration: read build file (yaml, xml)
Note right of integration: before install
integration->integration: apt-get install golang-go
Note right of integration: install
integration->integration: go get -u github.com/golang/dep/cmd/dep
Note right of integration: before script
integration->integration: download config file
integration->integration: set enviroment variables
integration->integration: dep ensure
Note right of integration: script
integration->integration: code convention (coding style)
integration->integration: code quality (DRY, Coupling, Dead Code)
integration->integration: security (white box)
integration->integration: unit test, coverage
integration->integration: build binary (build image)
integration->Remote Server: binary
Note right of Remote Server: API Test Server, Newman
Remote Server->Remote Server: database migration
Remote Server->Remote Server: insert test data
Remote Server->Remote Server: integration test (http request)
integration->Remote Server: binary
Note right of Remote Server: UAT Server, Selenium
Remote Server->Remote Server: database migration
Remote Server->Remote Server: insert test data
Remote Server->Remote Server: user acceptance testing
integration->Remote Server: binary
Note right of Remote Server: System Server, Selenium, Cloud Service
Remote Server->Remote Server: database migration
Remote Server->Remote Server: insert test data
Remote Server->Remote Server: system test (cross browser, device)
integration->Remote Server: binary
Note right of Remote Server: Performance Server, apache ab
Remote Server->Remote Server: database migration
Remote Server->Remote Server: insert test data
Remote Server->Remote Server: performance (stress test)
integration->Remote Server: binary
Note right of Remote Server: Security Server, HP WebInspect
Remote Server->Remote Server: database migration
Remote Server->Remote Server: insert test data
Remote Server->Remote Server: security testing (black box)
Note right of integration: after_success/after_failure
integration->integration: godoc
Note right of integration: after_failure
integration->integration: alert team
Note right of integration: before_deploy
integration->integration: build image
integration->integration: push image
Note right of integration: deploy
integration->integration: gcloud app deploy
integration->server: deploy
Note right of integration: after_deploy
integration->integration: check online server
Note right of integration: after_script
integration->integration: send notification
```

## Continuous Delivery

* Responding to change
  * Frequency Delivery

* Continuous
  * Break the bottleneck
  * Automation
  * Feedback
  * Correct

## Pipeline

![pipeline](https://jenkins.io/doc/book/resources/pipeline/realworld-pipeline-flow.png)

## Continuous Integration

* Reliability
* Testing
* Code Quality

## Testing

![test-quadrant](https://continuousdelivery.com/images/test-quadrant.png)

## Continuous Deploy

* Version to Environment
  * uat → sit
  * release → production
  * v1.2.0 → production
* Migration
  * Database
  * Configuration
* Zero Downtime
* Rollback

## Reference

* Continuous Delivery, Jez Humble, Devid Farley, 2014
* [Travis CI Docs](https://docs.travis-ci.com/user/customizing-the-build/#The-Build-Lifecycle)
* [Jenkins Handbook](https://jenkins.io/doc/book/)
* [Continuous delivery@wikipedia](https://en.wikipedia.org/wiki/Continuous_delivery)
* [架設CI Server(Jenkins)](http://joel-zhong.logdown.com/posts/545549/setup-ci-server-using-jenkins)