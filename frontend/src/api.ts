import axios from "axios";
import { 
  AttendeesListApi, 
  ProfileApi, 
} from "./schema";

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const api = axios.create({
  baseURL: baseURL,
});
const attendeesListApi = new AttendeesListApi(undefined, baseURL, api);
const profileApi = new ProfileApi(undefined, baseURL, api);

export {
  attendeesListApi,
  profileApi,
}
