function validateUname(uname) {
    if (uname && uname.length > 0) return true;
    return false;
}

function validatePswd(pswd) {
    if (pswd && pswd.length > 0) return true;
    return false;
}

new Vue({
    el: '#admin-app',
    data: function() {
        return {
            alertVisible: false,
            loginVisible: true,
            panelVisible: false,
            uname: "",
            pswd: "",
            alertContent: "",
            accessToken: undefined
        }
    },

    // define methods under the `methods` object
    methods: {
        signin: function(e) {
            if (!validateUname(this.uname)) {
                this.alertContent = "Invalid username";
                this.alertVisible = true;
                return;
            }
            if (!validatePswd(this.pswd)) {
                this.alertContent = "Invalid password";
                this.alertVisible = true;
                return;
            }
            $.ajax({
                url: "admin/login",
                type: 'POST',
                timeout: 8000,
                data: $('#login-form').serialize(),
                success: (function(resp) {
                    this.accessToken = resp;
                    this.loginVisible = false;
                }).bind(this),
                error: (function(err) {
                    this.alertContent = "Login error" + err;
                    this.alertVisible = true;
                }).bind(this)
            });
        }
    }
})