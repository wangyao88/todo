$(document).ready(function () {
    $("#submit").click(login);
});

function login() {
    var UserName = $("#UserName").val()
    var Password = $("#Password").val()

    $.ajax({
        type: "POST",
        url: "/login",
        data: {
            UserName: UserName,
            Password: Password
        },
        dataType: "json",
        success: function(data){
            if (data.status == 0) {
                alert(data.msg);
            } else{
                window.location.href = "/";
            }
        }
    });
}