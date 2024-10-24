import axios from "axios";
import {
  AttendeesListApi,
  ProfileApi,
  AccessHistoryApi,
  RankingApi,
  LAApi,
  SettingApi,
  AuthenticationApi,
} from "./schema";

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const api = axios.create({
  baseURL: baseURL,
});
const attendeesListApi = new AttendeesListApi(undefined, baseURL, api);
const profileApi = new ProfileApi(undefined, baseURL, api);
const accessHistoryApi = new AccessHistoryApi(undefined, baseURL, api);
const rankingApi = new RankingApi(undefined, baseURL, api);
const labAssistantApi = new LAApi(undefined, baseURL, api);
const settingApi = new SettingApi(undefined, baseURL, api);
const authenticationApi = new AuthenticationApi(undefined, baseURL, api);

export {
  attendeesListApi,
  profileApi,
  accessHistoryApi,
  rankingApi,
  labAssistantApi,
  settingApi,
  authenticationApi,
};
