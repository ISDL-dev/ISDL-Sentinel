import axios from "axios";
import { 
  PutStatusApi, 
  GetUserByIdApi, 
  PutAvatarApi 
} from "./schema";

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const api = axios.create({
  baseURL: baseURL,
});

const putStatusApi = new PutStatusApi(undefined, baseURL, api);
const getUserByIdApi = new GetUserByIdApi(undefined, baseURL, api);
const putAvatarApi = new PutAvatarApi(undefined, baseURL, api);

export {
  putStatusApi,
  getUserByIdApi,
  putAvatarApi
}
