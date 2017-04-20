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
            accessToken: "",
            activeTabIndex: 0,
            subscribers: [],
            subscribersSelection: [],
            feedbacks: []
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
                url: "admin/signin",
                type: 'POST',
                timeout: 8000,
                data: $('#login-form').serialize(),
                success: (function(resp) {
                    this.accessToken = resp;
                    this.loginVisible = false;
                }).bind(this),
                error: (function(err) {
                    this.alertContent = err;
                    this.alertVisible = true;
                }).bind(this)
            });
        },

        subTabClick: function(e) {
            this.activeTabIndex = 1;
            $.ajax({
                url: "admin/subscriber",
                type: 'GET',
                timeout: 8000,
                dataType: 'json',
                headers: { "Auth": this.accessToken },
                success: (function(resp) {
                    if (resp.data) {
                        this.subscribers = resp.data.map(function(subscriber) {
                            subscriber.active = subscriber.active.toString();
                            return subscriber;
                        });
                    } else {
                        this.alertContent = "Subscriber data unavailable";
                        this.alertVisible = true;
                    }
                }).bind(this),
                error: (function(err) {
                    this.loginVisible = true;
                    this.alertContent = err;
                    this.alertVisible = true;
                }).bind(this)
            });
        },

        fedTabClick: function(e) {
            this.activeTabIndex = 2;
            $.ajax({
                url: "admin/feedback",
                type: 'GET',
                timeout: 8000,
                dataType: 'json',
                headers: { "Auth": this.accessToken },
                success: (function(resp) {
                    if (resp.data) {
                        this.feedbacks = resp.data;
                    } else {
                        this.alertContent = "Feedback data unavailable";
                        this.alertVisible = true;
                    }
                }).bind(this),
                error: (function(err) {
                    this.loginVisible = true;
                    this.alertContent = err;
                    this.alertVisible = true;
                }).bind(this)
            });
        },

        handleSubscriberSelectionChange: function(val) {
            this.subscribersSelection = val;
        },

        replyFeedback: function(feedback) {
            window.open("mailto:feedback.email");
        }
    }
})