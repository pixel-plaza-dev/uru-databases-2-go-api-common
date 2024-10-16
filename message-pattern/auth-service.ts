// Auth service message patterns
const AUTH_PATTERN = {
    SIGN_UP: {cmd: 'sign_up'},
    LOG_IN: {cmd: 'log_in'},
    LOG_OUT: {cmd: 'log_out'},
    ADD_PERMISSION: {cmd: 'add_permission'},
    REMOVE_PERMISSION: {cmd: 'remove_permission'},
    ADD_ROLE_PERMISSION: {cmd: 'add_role_permission'},
    REMOVE_ROLE_PERMISSION: {cmd: 'remove_role_permission'},
    ADD_ROLE: {cmd: 'add_role'},
    REMOVE_ROLE: {cmd: 'remove_role'},
    REFRESH: {cmd: 'refresh'},
    CLOSE_SESSIONS: {cmd: 'close_sessions'},
    DELETE: {cmd: 'delete'},
}
export default AUTH_PATTERN
