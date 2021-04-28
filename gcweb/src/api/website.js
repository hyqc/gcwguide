import axios from "@/util/request.js";

const website = {
  add: "website/add",
  list: "website/list",
  delete: "website/delete",
  update: "website/update",
  groups: "website/groups",
  imageUpload: "upload/image",
};

export function WebList(data) {
  return axios({
    method: "get",
    data: data,
    url: website.list,
  });
}

export function WebAdd(data) {
  return axios({
    method: "post",
    data: data,
    url: website.add,
  });
}


export function WebEdit(data) {
  return axios({
    method: "post",
    data: data,
    url: website.update,
  });
}

export function WebDelete(data) {
  return axios({
    method: "post",
    data: data,
    url: website.delete,
  });
}

export function WebGroups(data) {
  return axios({
    method: "get",
    data: data,
    url: website.groups,
  });
}

export const WebLogoUploadUrl = website.imageUpload
