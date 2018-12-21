function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

!function () {
    laydate.skin('dahong');
    laydate({ elem: '#TodoExcpteEndTime' });
    laydate({ elem: '#TodoRealyEndTime' });
}();

$(document).ready(function() {
    $("#form_todo_add_save").click(saveOrUpdateTodo);
    if($("#TodoRealyEndTime").val() == "0001-01-01") {
        $("#TodoRealyEndTime").val("");
    }
});

function saveOrUpdateTodo() {
    if($("#TodoId").val()) {
        updateTodo();
    } else {
        saveTodo();
    }
}

function saveTodo() {
    $.ajax({
        type: "POST",
        url: "/todo/add",
        data: $("#form_todo").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("保存成功！！！");
                parent.getTodoTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}

function updateTodo() {
    $.ajax({
        type: "POST",
        url: "/todo/update",
        data: $("#form_todo").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("更新成功！！！");
                parent.getTodoTableData();
                TaskCancel();
            } else {
                alert("更新失败！！！")
            }
        }
    })
}

function TaskCancel() {
    var index = parent.layer.getFrameIndex(window.name)
    parent.layer.close(index);
}