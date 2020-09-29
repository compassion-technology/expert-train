$(document).ready(function(){
    $(".container-sub").hide();
      $(".media").click(function(){
          var name = 0;
        $(".message-container").empty();
        name = $('.media').attr('name');
        $.getJSON( "https://jsonplaceholder.typicode.com/comments", {postId:name},function( data ) {
        $.each( data, function( val ) {
            $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='../compassion-space-ui/images/profile.jpg' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>You</strong>This is a test " +data[val].email+ "</p></div>"); 
            //alert(data[val].email);
            //items.push( "<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='../compassion-space-ui/images/profile.jpg' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>You</strong>This is a test" +some.email[2]+ "</p></div>");
  });
          });
        $(".placeholder-container").remove();
        $(".container-sub").show();
      });
      $(".send").click(function(){
          var message = $.trim($(".message").val());
          $(".message").val('');
          $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='../compassion-space-ui/images/profile.jpg' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>You</strong>"+message+"</p></div>");
      });
    });