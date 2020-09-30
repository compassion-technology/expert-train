$(document).ready(function(){
    $(".container-sub").hide();
      $(".media").click(function(){
          var name = 0;
        $(".message-container").empty();
        name = $('.media').attr('name');
        $.getJSON( "http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/user/01234567/messages",function( data ) {
        $.each( data, function( val ) {
          var person = data[val].from;
          var from = "";
          var picture = "";
          if (person === "01234567")
          {
            from = "You";
            picture = "../images/profile.jpg";
          }
          else {
            from = data[val].from;
            picture = "../images/larpo.png";
          }
          if (data[val].message.content_url != null) {
            $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='"+picture+"' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong>"+from+"</strong><br>"+data[val].message.text.english+"<br><img class='mr-3 shadow-sm' src='"+data[val].message.content_url+"' width='100' height='100'></p></div>");
            console.log(data[val].message.content_url);
          }
          else{
            $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='"+picture+"' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>"+from+"</strong>" +data[val].message.text.english+ "</p></div>");            
          }
  });
          });
        $(".placeholder-container").remove();
        $(".container-sub").show();
      });
      $(".send").click(function(){
          var message = $.trim($(".message").val());
          $(".message").val('');
          $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='"+picture+"' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>"+from+"</strong>"+message+"</p></div>");
      });
    });