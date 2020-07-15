// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

def HEAD_COMMIT

pipeline {
    options {
        skipDefaultCheckout true
        disableConcurrentBuilds()
    }

    agent {
       docker {
            image "${RUNNER_DOCKER_IMAGE}"
            args "${RUNNER_DOCKER_ARGS}"
            registryUrl "${RUNNER_DOCKER_REGISTRY_URL}"
            registryCredentialsId 'ocir-pull-and-push-account'
        }
    }

    environment {
        DOCKER_CI_IMAGE_NAME = 'verrazzano-crd-generator-jenkins'
        DOCKER_PUBLISH_IMAGE_NAME = 'verrazzano-crd-generator-operator'
        DOCKER_IMAGE_NAME = "${env.BRANCH_NAME == 'master' ? env.DOCKER_PUBLISH_IMAGE_NAME : env.DOCKER_CI_IMAGE_NAME}"
        CREATE_LATEST_TAG = "${env.BRANCH_NAME == 'master' ? '1' : '0'}"
        GOPATH = '/home/opc/go'
        GO_REPO_PATH = "${GOPATH}/src/github.com/verrazzano"
        DOCKER_CREDS = credentials('ocir-pull-and-push-account')
    }

    stages {
        stage('Clean workspace and checkout') {
            steps {
                checkout scm
                sh """
                    echo "${DOCKER_CREDS_PSW}" | docker login ${env.DOCKER_REPO} -u ${DOCKER_CREDS_USR} --password-stdin
                    rm -rf ${GO_REPO_PATH}/verrazzano-crd-generator
                    mkdir -p ${GO_REPO_PATH}/verrazzano-crd-generator
                    tar cf - . | (cd ${GO_REPO_PATH}/verrazzano-crd-generator/ ; tar xf -)
                """
            }
        }

        stage('Build') {
            when { not { buildingTag() } }
            steps {
                sh """
                    cd ${GO_REPO_PATH}/verrazzano-crd-generator
		    make go-build 
                """
            }
        }

        stage('Copyright Compliance Check') {
            when { not { buildingTag() } }
            steps {
                copyrightScan "${WORKSPACE}"
            }
        }
    }

    post {
        failure {
            mail to: "${env.BUILD_NOTIFICATION_TO_EMAIL}", from: 'noreply@oracle.com',
            subject: "Verrazzano: ${env.JOB_NAME} - Failed",
            body: "Job Failed - \"${env.JOB_NAME}\" build: ${env.BUILD_NUMBER}\n\nView the log at:\n ${env.BUILD_URL}\n\nBlue Ocean:\n${env.RUN_DISPLAY_URL}"
        }
    }

}
