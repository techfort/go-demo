node {
    def commit_id
    stage('scm') {
        checkout scm
        sh 'git rev-parse --short HEAD > .git/commit-id'
        commit_id = readFile('.git/commit-id').trim()
    }

    stage('test') {
        sh 'go test ./...'
    }

    stage('docker build/publish' ) {
        docker.withRegistry('https://index.docker.io/v1/', 'dockerhub') {
            def app = docker.build("techfort/go-demo:${commit_id}", '.').push()
        }
    }
}node {
    def commit_id
    
    def root = tool name: 'golang1.13', type: 'go'
 
    // Export environment variables pointing to the directory where Go was installed
    
    stage('scm') {
        checkout scm
        sh 'git rev-parse --short HEAD > .git/commit-id'
        commit_id = readFile('.git/commit-id').trim()
    }

    stage('test') {
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "CGO_ENABLED=0"]) {
            sh 'go test ./...'
        }
    }

    stage('docker build/publish') {
        docker.withRegistry('https://index.docker.io/v1/', 'dockerhub') {
            def app = docker.build("techfort/go-demo:${commit_id}", '.').push()
        }
    }
}