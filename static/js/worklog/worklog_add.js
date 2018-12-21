function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

!function () {
    laydate.skin('dahong');
    laydate({ elem: '#WorkLogCreateTime' });
    laydate({ elem: '#WorkLogUpdateTime' });
}();

$(document).ready(function() {
    $("#form_workLog_add_save").click(saveOrUpdateWorkLog);
    var projectId = $("#projectId").val();
    $("#selectProjectId option[value='"+projectId+"']").attr("selected",true);
});

function saveOrUpdateWorkLog() {
    if($("#WorkLogId").val()) {
        updateWorkLog();
    } else {
        saveWorkLog();
    }
}

function saveWorkLog() {
    $.ajax({
        type: "POST",
        url: "/workLog/add",
        data: $("#form_workLog").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("保存成功！！！");
                parent.getWorkLogTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}

function updateWorkLog() {
    $.ajax({
        type: "POST",
        url: "/workLog/update",
        data: $("#form_workLog").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("更新成功！！！");
                parent.getWorkLogTableData();
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