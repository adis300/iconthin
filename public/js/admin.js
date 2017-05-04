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
            feedbacks: [],
            formatDate: function(timestamp) {
                var date = new Date(timestamp * 1000);
                return date.toLocaleString();
            },
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
            var subject = "RE:%20EMPTY SUBJECT";
            if (feedback.subject) subject = "RE:%20" + encodeURIComponent(feedback.subject);

            var body = "\n\n";
            if (feedback.name || feedback.email) {
                body += "From: ";
                if (feedback.name) {
                    body += feedback.name;
                    if (feedback.email) body += "(" + feedback.email + ")";
                } else {
                    body += feedback.email
                }
                body += "\n";
            }
            if (feedback.subject)
                body += "Subject: " + feedback.subject + "\n";
            if (feedback.timestamp)
                body += "Date: " + this.formatDate(feedback.timestamp) + "\n";

            if (feedback.phone)
                body += "Phone: " + feedback.phone + "\n";
            if (feedback.body)
                body += "\n" + feedback.body + "\n";

            body = encodeURIComponent(body);
            window.open("mailto:" + feedback.email + "?subject=" + subject + "&body=" + body);
        }
    }
})