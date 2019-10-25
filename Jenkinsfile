node {
    def commit_id
    
    def root = tool name: 'golang1.13', type: 'go'
    def to = emailextrecipients([
          [$class: 'CulpritsRecipientProvider'],
          [$class: 'DevelopersRecipientProvider'],
          [$class: 'RequesterRecipientProvider']
    ])
    // Export environment variables pointing to the directory where Go was installed
    try {   
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
    } catch(e) {
        // mark build as failed
        currentBuild.result = "FAILURE";
        // set variables
        def subject = "${env.JOB_NAME} - Build #${env.BUILD_NUMBER} ${currentBuild.result}"
        def content = '${JELLY_SCRIPT,template="html"}'

        // send email
        if(to != null && !to.isEmpty()) {
        emailext(body: content, mimeType: 'text/html',
            replyTo: '$DEFAULT_REPLYTO', subject: subject,
            to: to, attachLog: true )
        }

        // mark current build as a failure and throw the error
        throw e;
    }
}