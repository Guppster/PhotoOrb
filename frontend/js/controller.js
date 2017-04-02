var three60Controller = angular.module('three60Controller', [])

three60Controller.controller('Three60Controller', ['$scope', '$http', '$timeout',
  function ($scope, $http, $timeout) {

    if (window.JpegCamera) {
      var camera;
      $scope.signedInUser = user;

      var itemId = itemId ? itemId : 10;
      var getNewId;

      var take_snapshots = function (count, id) {
        var snapshot = camera.capture();

        if (JpegCamera.canvas_supported()) {
          snapshot.get_canvas(add_snapshot);
        }
        else {
          var image = document.createElement("img");
          image.src = "no_canvas_photo.jpg";
          setTimeout(function () {
            add_snapshot.call(snapshot, image)
          }, 1);
        }


        if (count > 1) {
          setTimeout(function () {
            take_snapshots(count - 1);
          }, 500);
        }
      };

      var add_snapshot = function (element) {
        $(element).data("snapshot", this).addClass("item");
        console.log($(element).data("snapshot")._extra_canvas)

        var itemToUpload = $(element).data("snapshot");
        upload_snapshot(itemToUpload);

        var $container = $("#snapshots").append(element);
        var $camera = $("#camera");
        var camera_ratio = $camera.innerWidth() / $camera.innerHeight();

        var height = $container.height()
        element.style.height = "" + height + "px";
        element.style.width = "" + Math.round(camera_ratio * height) + "px";

        var scroll = $container[0].scrollWidth - $container.innerWidth();

        $container.animate({
          scrollLeft: scroll
        }, 200);
      };

      var select_snapshot = function () {
        $(".item").removeClass("selected");
        var snapshot = $(this).addClass("selected").data("snapshot");
        $("#discard_snapshot, #upload_snapshot, #api_url").show();
        snapshot.show();
        $("#show_stream").show();
      };

      var clear_upload_data = function () {
        $("#upload_status, #upload_result").html("");
      };

      var upload_snapshot = function (item) {
        if (!apiUrl.length) {
          $("#upload_status").html("Please provide URL for the upload");
          return;
        }

        clear_upload_data();
        $("#loader").show();
        $("#upload_snapshot").prop("disabled", true);

        var snapshot = item ? item : $(".item.selected").data("snapshot");


        console.log(snapshot);
        console.log(snapshot._extra_canvas);

        itemId += 1;
        console.log(itemId);

        if (snapshot._extra_canvas.toBlob) {
          snapshot._extra_canvas.toBlob(
            function (blob) {
              var formData = new FormData();
              formData.append('uploadfile', blob, itemId + ".jpg");
              var seriesPath = apiUrl + "/" + $scope.signedInUser + "/" + $scope.userID;
              console.log($scope.userID);
              $.ajax({
                type: "POST",
                url: seriesPath,
                data: formData,
                processData: false,
                contentType: false,
              }).done(function () {
                upload_done();
              }).fail(function () {
                upload_fail();
              }).always(function () {
                // alert("complete");
              });
            },
            'image/jpg'
          );
        }
      };

      var upload_done = function (response) {
        $("#upload_snapshot").prop("disabled", false);
        $("#loader").hide();
        $("#upload_status").html("<span style='font-size: 30px' class='glyphicon glyphicon-ok'></span>");
        $("#upload_result").html(response);
      };

      var upload_fail = function (code, error, response) {
        $("#upload_snapshot").prop("disabled", false);
        $("#loader").hide();
        $("#upload_status").html(
          "Upload failed with status " + code + " (" + error + ")");
        $("#upload_result").html(response);
      };

      var discard_snapshot = function () {
        var element = $(".item.selected").removeClass("item selected");

        var next = element.nextAll(".item").first();

        if (!next.size()) {
          next = element.prevAll(".item").first();
        }

        if (next.size()) {
          next.addClass("selected");
          next.data("snapshot").show();
        }
        else {
          hide_snapshot_controls();
        }

        element.data("snapshot").discard();

        element.hide("slow", function () {
          $(this).remove()
        });
      };

      var show_stream = function () {
        $(this).hide();
        $(".item").removeClass("selected");
        hide_snapshot_controls();
        clear_upload_data();
        camera.show_stream();
      };

      var hide_snapshot_controls = function () {
        $("#discard_snapshot, #upload_snapshot, #api_url").hide();
        $("#upload_result, #upload_status").html("");
        $("#show_stream").hide();
      };


      $("#take_snapshots").click(function () {
        // itemId = 10;

        if ($scope.signedInUser) {
          getNewId = $.ajax(
            {
              type: "POST",
              url: apiUrl + "/" + $scope.signedInUser,
              data: {},
              processData: false,
              contentType: false,
            }
          );
        }
        getNewId.done(function (id) {
          console.log(id);
          $scope.userID = id;
          take_snapshots(20, id);
          $scope.showPreview = true;
        });
      });

      $("#snapshots").on("click", ".item", select_snapshot);
      $("#upload_snapshot").click(upload_snapshot);
      $("#discard_snapshot").click(discard_snapshot);
      $("#show_stream").click(show_stream);

      var options = {
        shutter_ogg_url: "../dist/shutter.ogg",
        shutter_mp3_url: "../dist/shutter.mp3",
        swf_url: "../dist/jpeg_camera.swf"
      }

      camera = new JpegCamera("#camera", options).ready(function (info) {
        $("#take_snapshots").show();

        $("#camera_info").html(
          "Camera resolution: " + info.video_width + "x" + info.video_height);
      });
    }
  }
])
;


const apiUrl = 'http://ec2-54-175-181-19.compute-1.amazonaws.com/upload';
const user = "nick"
