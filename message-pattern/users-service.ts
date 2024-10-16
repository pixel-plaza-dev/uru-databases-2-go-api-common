// Users service message patterns
const USERS_PATTERN = {
    UPDATE_USER: {cmd: 'update_user'},
    CHANGE_USERNAME: {cmd: 'change_username'},
    CHANGE_PASSWORD: {cmd: 'change_password'},
    CHANGE_EMAIL: {cmd: 'change_email'},
    SEND_EMAIL_VERIFICATION_TOKEN: {cmd: 'send_email_verification_token'},
    VERIFY_EMAIL: {cmd: 'verify_email'},
    FORGOT_PASSWORD: {cmd: 'forgot_password'},
    RESET_PASSWORD: {cmd: 'reset_password'},
}
export default USERS_PATTERN;