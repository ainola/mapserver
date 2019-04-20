/* exported SearchResult */
/* globals SearchStore: true */
/* globals layerMgr: true */

var SearchResult = {
  view: function(vnode){
    var map = vnode.attrs.map;

    function getLayer(obj){
      var layer = layerMgr.getLayerByY(obj.y);
      return layer ? layer.name : "<unknown>";
    }

    function getPos(obj){
      var layer = layerMgr.getLayerByY(obj.y);
      var text = obj.x + "/" + obj.y + "/" + obj.z;

      return m("span", {class:"badge badge-success"}, text);
    }

    var rows = SearchStore.result.map(function(obj){

      var row_classes = "";
      var description = obj.type;
      var type = obj.type;

      if (obj.type == "poi"){
        description = m("span", obj.attributes.name);
        type = m("img", { src: "css/images/marker-icon.png" });
      }

      if (obj.type == "shop") {
        if (obj.attributes.stock == 0){
          row_classes += "table-warning";
          type = m("img", { src: "pics/shop_empty.png" });
        } else {
          type = m("img", { src: "pics/shop.png" });
        }

        description = m("span", [
          "Shop, trading ",
          m("span", {class:"badge badge-primary"}, obj.attributes.out_count + "x"),
          m("span", {class:"badge badge-info"}, obj.attributes.out_item),
          " for ",
          m("span", {class:"badge badge-primary"},  obj.attributes.in_count + "x"),
          m("span", {class:"badge badge-info"},  obj.attributes.in_item),
          " Stock: ",
          m("span", {class:"badge badge-info"}, obj.attributes.stock)
        ]);
      }

      function onclick(){
        map.setView([obj.z, obj.x], 12);
      }

      return m("tr", {"class": row_classes}, [
        m("td", type),
        m("td", obj.attributes.owner),
        m("td", getLayer(obj)),
        m("td", getPos(obj)),
        m("td", description),
        m("button[type=button]", {class: "btn btn-secondary", onclick: onclick }, [
          "Goto ",
          m("i", { class: "fas fa-play" })
        ])
      ]);
    });

    return m("table", {class:"table table-striped"}, [
      m("thead", [
        m("tr", [
          m("th", "Type"),
          m("th", "Owner"),
          m("th", "Layer"),
          m("th", "Position"),
          m("th", "Description"),
          m("th", "Action")
        ])
      ]),
      m("tbody", rows)
    ]);
  }
};
