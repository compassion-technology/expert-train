$(document).ready(function () {
  $(".container-sub").hide();
  $(".media").click(function () {
    var clicked = $(this).attr('id');
    if (clicked == "123"){
      $(".etonome").removeClass("highlight");
      $(".larpo").addClass("highlight");
    }
    if (clicked == "456"){
      $(".larpo").removeClass("highlight");
      $(".etonome").addClass("highlight");
    }
    $(".message-container").empty();
    groupId = parseInt($(this).attr('id'));
    getMessages(groupId);
    $(".placeholder-container").remove();
    $(".container-sub").show();
  });

  $(".view").click(function () {
    
    var type = $(this).attr('id');
    if (type == "sponsor") {
        $(".sponsor").remove();
        $(".beneficiary").remove();
      $(".nav-underline").append("<a class='nav-link beneficiary' href='#'>Sponsored Child/Children<span class='badge badge-pill bg-light align-text-bottom'>2</span></a>");
    }
    else {
        $(".beneficiary").remove();
        $(".sponsor").remove();
      $(".nav-underline").append("<a class='nav-link sponsor' href='letters.html'>Sponsor Profile</a>");

    }
  });

  $(".send").click(function () {
    var text = document.getElementById('message-area').innerText;
    var imagepath = "";
    if ($("#image-container").length != 0){
      imagepath = document.getElementById('image-container').innerText;
    }
    postMessages(groupId,text,imagepath);
    document.getElementById('message-area').innerText = "";
    getLastMessage(groupId,text, imagepath);
    $(".image-container").remove();
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
      $(".text-container").append("<div contenteditable='false' id='image-container' class='border border-gray image-container' style='width:100%; padding:1em'>"+filename+"</div>");
    } 
  });

  //here's the code for record
let preview = document.getElementById("preview");
let recording = document.getElementById("recording");
let startButton = document.getElementById("startButton");
let stopButton = document.getElementById("stopButton");
let downloadButton = document.getElementById("downloadButton");
let logElement = document.getElementById("log");

let recordingTimeMS = 5000;
function log(msg) {
  logElement.innerHTML += msg + "\n";
}
function wait(delayInMS) {
  return new Promise(resolve => setTimeout(resolve, delayInMS));
}
function startRecording(stream, lengthInMS) {
  let recorder = new MediaRecorder(stream);
  let data = [];
 
  recorder.ondataavailable = event => data.push(event.data);
  recorder.start();
  log(recorder.state + " for " + (lengthInMS/1000) + " seconds...");
 
  let stopped = new Promise((resolve, reject) => {
    recorder.onstop = resolve;
    recorder.onerror = event => reject(event.name);
  });

  let recorded = wait(lengthInMS).then(
    () => recorder.state == "recording" && recorder.stop()
  );
 
  return Promise.all([
    stopped,
    recorded
  ])
  .then(() => data);
}
function stop(stream) {
  stream.getTracks().forEach(track => track.stop());
}
startButton.addEventListener("click", function() {
  navigator.mediaDevices.getUserMedia({
    video: true,
    audio: true
  }).then(stream => {
    preview.srcObject = stream;
    downloadButton.href = stream;
    preview.captureStream = preview.captureStream || preview.mozCaptureStream;
    return new Promise(resolve => preview.onplaying = resolve);
  }).then(() => startRecording(preview.captureStream(), recordingTimeMS))
  .then (recordedChunks => {
    let recordedBlob = new Blob(recordedChunks, { type: "video/webm" });
    recording.src = URL.createObjectURL(recordedBlob);
    downloadButton.href = recording.src;
    downloadButton.download = "RecordedVideo.webm";
    
    log("Successfully recorded " + recordedBlob.size + " bytes of " +
        recordedBlob.type + " media.");
  })
  .catch(log);
}, false);stopButton.addEventListener("click", function() {
  stop(preview.srcObject);
}, false);
});

function getLastMessage(groupId, text, imagepath) {

      // var from = "You";
      // var picture = "../compassion-space-ui/images/profile.jpg";
      // if(imagepath){
      //   $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong>" + from + "</strong><br>" + text + "<br><img class='mr-3 shadow-sm' src='" + imagepath + "' width='100' height='100'></p></div>");
      // }
      // else {
      //   $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>" + from + "</strong>" + text + "</p></div>");
      // }
}

function getMessages(groupId) {
  $.getJSON("http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/messages?take=5&group="+groupId+"", function (data) {
    data.sort(function(a,b) {
      var acreated = new Date(a.created);
      var bcreated = new Date(b.created);
      return (acreated - bcreated);
    });
    $.each(data, function (val) {
      var person = data[val].from;
      var from = "";
      var picture = "";
      if (person === "01234567") {
        from = "You";
        picture = "https://previews.123rf.com/images/arekmalang/arekmalang1207/arekmalang120700008/14383125-a-shot-of-a-smiling-confident-asian-indian-businesswoman-outdoor.jpg";
      }
      else {
        from = data[val].from;
        picture = "https://us.123rf.com/450wm/borgogniels/borgogniels1802/borgogniels180200002/95510792-little-native-african-boy-standing-outdoors-under-the-rain.jpg?ver=6";
      }
      var content_url = data[val].content_url;
      if (content_url) {
        $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong>" + from + "</strong><br>" + data[val].message.en + "<br><img class='mr-3 shadow-sm' src='" + data[val].content_url + "' width='100' height='100'></p></div>");
      }
      else {
        $(".message-container").append("<div class='media text-muted pt-3'><img class='mr-3 rounded shadow-sm' src='" + picture + "' width='32' height='32'><p class='media-body pb-3 mb-0 small lh-125 border-bottom border-gray'><strong class='d-block text-gray-dark'>" + from + "</strong>" + data[val].message.en + "</p></div>");
      }
    });
  });
}

async function postMessages(groupId, text, imagepath) {
  var actual = "";
  var type = "";
  if (imagepath != ""){
    var encodedUrl = await getBase64Image();
    var commaSeparated = encodedUrl.split(',');
    actual = commaSeparated[commaSeparated.length - 1]
    type = "jpg";
  }
  else {
    actual = "";
    type = "";
  }

    var someData = {
        "group_id":groupId,
        "text": {
            "source": text
        },
        "content": {
          "type": type,
          "data": actual
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

const getBase64Image=() => new Promise(resolve => {
  var filesSelected = document.getElementById("input-file-now").files;
  var fileToLoad = filesSelected[0];
  const reader = new window.FileReader()
  reader.onload = () => {
    resolve(reader.result)
  }
  reader.onerror = e => {
    console.error(e)
    window.alert(UPLOAD_ERROR_MESSAGE)
    metrics.track('File Upload Error', { Message: e.message, Alert: UPLOAD_ERROR_MESSAGE })
    resolve(null)
  }
  reader.readAsDataURL(fileToLoad)
})