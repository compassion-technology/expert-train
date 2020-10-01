$(document).ready(function () {
  $(".container-sub").hide();
  $(".media").click(function () {
    $(".message-container").empty();
    groupId = parseInt($(this).attr('id'));
    getMessages(groupId);
    $(".placeholder-container").remove();
    $(".container-sub").show();
  });

  $(".send").click(function () {
    var text = document.getElementById('message-area').innerText;
    postMessages(groupId,text);
    document.getElementById('message-area').innerText = "";
    getLastMessage(groupId);
  });

  $(".post").click(function () {
    var url = $(".file-upload").val();
    if (url == null || url=='')
    {
      $(".file-upload").append("<p>You must select a file to post</p>")
    }
    else
    {
      var filename = url.split("\\").pop();
      $(".text-container").append("<div contenteditable='false' class='border border-gray image-container' style='width:100%; padding:1em'>"+filename+"</div>");
    } 
  });
});

function getLastMessage(groupId) {
  $.getJSON("http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/messages?take=1&group="+groupId+"", function (data) {
    $.each(data.sort, function (val) {
      var person = data[val].from;
      var from = "You";
      var picture = "../compassion-space-ui/images/profile.jpg";
      $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>" + from + "</strong>" + data[val].message.text.english + "</p></div>");
    });
  });
}

function getMessages(groupId) {
  $.getJSON("http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/messages?take=5&group="+groupId+"", function (data) {
    $.each(data, function (val) {
      var person = data[val].from;
      var from = "";
      var picture = "";
      if (person === "01234567") {
        from = "You";
        picture = "../compassion-space-ui/images/profile.jpg";
      }
      else {
        from = data[val].from;
        picture = "../compassion-space-ui/images/larpo.png";
      }
      if (data[val].message.content_url != null) {
        $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong>" + from + "</strong><br>" + data[val].message.text.english + "<br><img class='mr-3 shadow-sm' src='" + data[val].message.content_url + "' width='100' height='100'></p></div>");
        console.log(data[val].message.content_url);
      }
      else {
        $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>" + from + "</strong>" + data[val].message.text.english + "</p></div>");
      }
    });
  });
}

function postMessages(groupId, text) {
    var someData = {
        "group_id":groupId,
        "text": {
            "amharic": "ከእርስዎ ለመስማት በጣም ጥሩ ነው! ከሚስትዎ ቀዶ ጥገና ጋር ሁሉም ነገር በጥሩ ሁኔታ እየጸለይኩ ነው!",
            "english": text
        }
      }
    var stringyfiedData = JSON.stringify(someData)
    var saveData = $.ajax({
      type: 'POST',
      contentType: "application/json",
      url: "http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/user/01234567/message",
      data: stringyfiedData,
      success: function(resultData) { 
        console.log("save complete"); 
      }
    });
}