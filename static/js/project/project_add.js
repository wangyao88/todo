function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

!function () {
    laydate.skin('dahong');
    laydate({ elem: '#ProjectStartTime' });
    laydate({ elem: '#ProjectEndTime' });
}();

$(document).ready(function() {
    $("#form_project_add_save").click(saveOrUpdateProject);
    var companyId = $("#companyId").val();
    $("#selectCompanyId option[value='"+companyId+"']").attr("selected",true);
});

function saveOrUpdateProject() {
    if($("#ProjectId").val()) {
        updateProject();
    } else {
        saveProject();
    }
}

function saveProject() {
    $.ajax({
        type: "POST",
        url: "/project/add",
        data: $("#form_project").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("保存成功！！！");
                parent.getProjectTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}

function updateProject() {
    $.ajax({
        type: "POST",
        url: "/project/update",
        data: $("#form_project").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("更新成功！！！");
                parent.getProjectTableData();
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