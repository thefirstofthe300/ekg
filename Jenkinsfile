pipeline {
  agent any
  options {
    checkoutToSubdirectory('src/github.com/thefirstofthe300/ekg')
  }
  tools {
    go 'Go 1.8'
  }
  environment {
    GOPATH = "$WORKSPACE"
    PATH = "$PATH:$GOPATH/bin"
  }
  stages {
    stage('Run unit tests') {
      steps {
        sh '''
go get -u github.com/golang/dep/cmd/dep
cd src/github.com/thefirstofthe300/ekg
dep ensure
go test $(go list ./... | grep -v /vendor/)
'''
      }
    }
    stage('Build') {
      steps {
        sh 'go build'
      }
    }
    stage('Deploy') {
      steps {
        echo 'Deploying'
      }
    }
  }
}

