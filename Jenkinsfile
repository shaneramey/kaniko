#! /usr/bin/env groovy

node() {
    stage('Checkout') {
        checkout scm
    }
    stage('Test') {
        sh "make test"
    }
}

