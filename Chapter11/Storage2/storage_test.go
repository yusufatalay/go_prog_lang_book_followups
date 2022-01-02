package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifier(t *testing.T) {

	// Save and restore original notifyUser

	saved := notifyUser

	defer func() { notifyUser = saved }()

	// Install the test's fake notifyUser

	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	// simulate 980MB used cond
	const user = "joe@example.com"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"

	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>,"+"want substring %q", notifiedMsg, wantSubstring)
	}

}
