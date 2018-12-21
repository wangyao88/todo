function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

!function () {
    laydate.skin('dahong');
    laydate({ elem: '#LeaveStartTime' });
    laydate({ elem: '#LeaveEndTime' });
}();

$(document).ready(function() {
    $("#form_leave_add_save").click(saveOrUpdateLeave);
    var leaveType = $("#leaveType").val();
    $("#selectLeaveType option[value='"+leaveType+"']").attr("selected",true);
});

function saveOrUpdateLeave() {
    if($("#LeaveId").val()) {
        updateLeave();
    } else {
        saveLeave();
    }
}

function saveLeave() {
    $.ajax({
        type: "POST",
        url: "/leave/add",
        data: $("#form_leave").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("保存成功！！！");
                parent.getLeaveTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}

function updateLeave() {
    $.ajax({
        type: "POST",
        url: "/leave/update",
        data: $("#form_leave").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("更新成功！！！");
                parent.getLeaveTableData();
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