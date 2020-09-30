$(document).ready(function () {
  $(".container-sub").hide();
  $(".media").click(function () {
    var name = 0;
    $(".message-container").empty();
    name = $('.media').attr('name');
    $.getJSON("http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/user/01234567/messages", function (data) {
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
    $(".placeholder-container").remove();
    $(".container-sub").show();
  });
  $(".send").click(function () {

    $.post( "https://jsonplaceholder.typicode.com/posts", { 
      "group_id":123,
      "text": {
          "amharic": "ከእርስዎ ለመስማት በጣም ጥሩ ነው! ከሚስትዎ ቀዶ ጥገና ጋር ሁሉም ነገር በጥሩ ሁኔታ እየጸለይኩ ነው!",
          "english": "Priya's awesome test"
      },
      "content": {
          "type": "jpg",
          "data": "{{base64 encoded file bytes}}"
      }
    })
    .done(function( data ) {
      alert( "Data Loaded: " + data );
    });
    // var message = $.trim($(".message").val());
    // $(".message").val('');
    // $(".image-conatiner").remove();

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