@Library('shared-library') _
String manifestName = null
node{
    properties([
        parameters([
            string(name: 'BRANCH_NAME', defaultValue: 'master', description: 'Branch to use'),
            string(name: 'GIT_URL', description: 'Source repo url'),
            string(name: 'GIT_AUTH_CREDS_ID', description: 'Credential id for git auth'),
            string(name: 'REGISTRY_AUTH_CREDS_ID', description: 'Credential id for registry auth'),
            string(name: 'REGISTRY', defaultValue: 'docker.io', description: 'Registry which is used to push docker images'),
            string(name: 'REPOSITORY',  description: 'Repository which is used to store images'),
            string(name: 'REGISTRY_USER', description: 'User name to login to registry'),
            string(name: 'TAG', defaultValue: 'test', description: 'Tag for docker images'),
            booleanParam(name: 'PUSH', defaultValue: false, description: 'Push to registry or not'),
        ])
    ])
    stage('Checkout repo') {
        git_checkout(env.BRANCH_NAME, env.GIT_AUTH_CREDS_ID, env.GIT_URL)
    }
    stage('Build docker image') {
        dir('repo') {
                manifestName = UUID.randomUUID().toString()
                Map<String, String> map = [
                    'context': '.',
                    'manifestName': manifestName
                ]
                for (platform in ['linux/amd64', 'linux/386']) {
                    map['platform'] = platform
                    build_image(map)
                }
        }
    }

    stage('Push docker image to docker hub') {
    }
        dir('repo') {
            if (params.PUSH) {
                push_image([
                    'registryAuthCredsId': env.REGISTRY_AUTH_CREDS_ID,
                    'registry': env.REGISTRY,
                    'registryUser': env.REGISTRY_USER,
                    'repository': env.REPOSITORY,
                    'tag': env.TAG,
                    'manifestName': manifestName])
            }
        }
}
