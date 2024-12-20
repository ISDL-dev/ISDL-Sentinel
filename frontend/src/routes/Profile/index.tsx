import { Box, Grid } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { profileApi } from "../../api";
import Banner from "../../features/Banner";
import AvatarList from "../../features/AvatarList";
import { GetUserById200Response, Avatar } from "../../schema";
import { useUser } from '../../userContext';

export default function Profile() {
  const [userData, setUserData] = useState<GetUserById200Response | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);

  const location = useLocation();
  const userId = location.state.userId;
  const { authUser, setAuthUser } = useUser();  
  const authUserId = authUser ? authUser.user_id : undefined;

  useEffect(() => {
    async function fetchUserData() {
      try {
        const response = await profileApi.getUserById(userId);
        setUserData(response.data);
      } catch (err) {
        setError("データの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    }

    fetchUserData();
  }, [userId]);

  const updateUserData = async (userId: number, avatarId: number) => {
    if (!authUser) return;
    try {
      if (userId === authUserId) {
        const requestBody: Avatar = {
          user_id: userId,
          avatar_id: avatarId,
        };
        await profileApi.putAvatar(requestBody);
        const response = await profileApi.getUserById(userId); // 再取得
        setUserData(response.data);
        setAuthUser({
          ...authUser,
          avatar_id: response.data.avatar_id,
          avatar_img_path: response.data.avatar_img_path
        });
      }
    } catch (err) {
      setError("アバターの更新に失敗しました");
    }
  };

  const deleteUserAvatar = async (userId: number, avatarId: number) => {
    if (!authUser) return;
    try {
      if (userId === authUserId) {
        const requestBody: Avatar = {
          user_id: userId,
          avatar_id: avatarId,
        };
        await profileApi.deleteAvatar(requestBody);
        const response = await profileApi.getUserById(userId); // 再取得
        setUserData(response.data);
        setAuthUser({
          ...authUser,
          avatar_id: response.data.avatar_id,
          avatar_img_path: response.data.avatar_img_path
        });
      }
    } catch (err) {
      setError("アバターの削除に失敗しました");
    }
  };

  const handleFileUpload = async () => {
    if (!selectedFile || !authUser) return;

    try {
      const response = await profileApi.postAvatar(userId, selectedFile);
      console.log(response);
      // Handle success (e.g., refetch avatar list)
      const updatedData = await profileApi.getUserById(userId);
      setUserData(updatedData.data);
      setAuthUser({
        ...authUser,
        avatar_id: updatedData.data.avatar_id,
        avatar_img_path: updatedData.data.avatar_img_path
      });
      setSelectedFile(null);
    } catch (error) {
      setError("アバターのアップロードに失敗しました");
    }
  };

  const handleCancelUpload = () => {
    setSelectedFile(null);
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error}</p>;

  if (!userData) return null;

  return (
    <Box pt={{ base: "80px", md: "80px", xl: "10px" }} pb={{ base: "80px", md: "80px", xl: "10px" }}>
      {/* Main Fields */}
      <Grid
        templateColumns={{
          base: "1fr",
          lg: "1.62fr 1fr",
        }}
        templateRows={{
          base: "repeat(2, 1fr)",
          lg: "1fr",
        }}
        gap={{ base: "20px", xl: "20px" }}
      >
        <Banner
          gridArea="1 / 1 / 2 / 2"
          banner="rgba(79, 209, 197, 1)"
          avatar={userData.avatar_img_path}
          name={userData.user_name}
          attendance_days={userData.attendance_days}
          stay_time={userData.stay_time}
          number_of_coin={userData.number_of_coin}
          email={userData.mail_address}
          grade={userData.grade}
          roleList={userData.role_list}
        />
        <AvatarList
          gridArea={{ base: "2 / 1 / 3 / 2", lg: "1 / 2 / 2 / 3" }}
          avatars={userData.avatar_list}
          onAvatarClick={(avatarId) => updateUserData(userData.user_id, avatarId)}
          onDeleteClick={(avatarId) => deleteUserAvatar(userData.user_id, avatarId)}
          selectedFile={selectedFile}
          setSelectedFile={setSelectedFile}
          onSaveUpload={handleFileUpload}
          onCancelUpload={handleCancelUpload}
        />
      </Grid>
    </Box>
  );
}
