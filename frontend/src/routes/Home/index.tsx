import {
  Avatar,
  Button,
  Flex,
  Grid,
  Spinner,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import "./Home.css";
import dayjs from "dayjs";
import "dayjs/locale/ja";
import { inRoom, overnight } from "../../models/users/user";
import { attendeesListApi } from "../../api";
import { GetAttendeesList200ResponseInner } from "../../schema";
import { useEffect, useState } from "react";
import { useUser } from "../../userContext";
import { useNavigate } from "react-router-dom";
import { Loading } from "../../features/Loading/Loading";

dayjs.locale("ja");

const decodeDate = (dateString: string) => {
  const date = dayjs(dateString);
  return `${dayjs(date).format("MM月DD日")}（${dayjs(date).format(
    "ddd"
  )}）${dayjs(date).format("HH時mm分")}`;
};

const isBetween8PMandMidnight = () => {
  const currentTime = dayjs();
  const hour = currentTime.hour();
  return hour >= 13 && hour < 24; // 20時以降の条件
};

function Home() {
  const { authUser, setAuthUser } = useUser();
  const [isFetching, setIsFetching] = useState<boolean>(false);
  const [attendeeList, setAttendeeList] = useState<
    GetAttendeesList200ResponseInner[] | null
  >(null);
  const navigate = useNavigate();

  const fetchAttendeesList = async () => {
    try {
      setIsFetching(true);
      const response = await attendeesListApi.getAttendeesList();
      setAttendeeList(response.data);
      setIsFetching(false);
    } catch (error) {
      console.log(error);
    }
  };

  const handleStatusChange = async () => {
    if (!authUser) return;
    try {
      const user = await attendeesListApi.putStatus({
        user_id: authUser.user_id,
        status: authUser.status,
      });
      setAuthUser({
        ...authUser,
        status: user.data.status,
      });
      await fetchAttendeesList(); // ここでリストを再取得
    } catch (error) {
      console.log(error);
    }
  };

  const handleStatusOvernightChange = async () => {
    if (!authUser) return;
    try {
      const user = await attendeesListApi.putStatus({
        user_id: authUser.user_id,
        status: overnight,
      });
      setAuthUser({
        ...authUser,
        status: user.data.status,
      });
      await fetchAttendeesList();
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    fetchAttendeesList();
  }, []);

  return (
    <div>
      <Grid
        templateColumns="repeat(3, 1fr)"
        alignItems={"center"}
        w={"-moz-max-content"}
        column={3}
      >
        <h1 className="block mb-1 text-4xl font-bold text-gray-900 dark:text-white p-3 text-left">
          出席者一覧
        </h1>
        {authUser && (
          <Grid templateColumns="repeat(3, 1fr)" alignItems={"center"}>
          {authUser.status === overnight ? (
            <>
              <Button
                colorScheme="cyan"
                variant="solid"
                size="lg"
                width={36}
                isDisabled={true}
              >
                宿泊済
              </Button>
              <Flex ml="auto">
                <Button
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  width={36}
                  isDisabled={true}
                >
                  入室済
                </Button>
                <Button
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  width={36}
                  onClick={handleStatusChange}
                >
                  退室
                </Button>
              </Flex>
            </>
          ) : authUser.status === inRoom ? (
            <>
              {isBetween8PMandMidnight() && (
                <Button
                  colorScheme="cyan"
                  variant="solid"
                  size="lg"
                  width={36}
                  onClick={handleStatusOvernightChange}
                >
                  宿泊
                </Button>
              )}
              <Flex ml="auto">
                <Button
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  width={36}
                  isDisabled={true}
                >
                  入室済
                </Button>
                <Button
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  width={36}
                  onClick={handleStatusChange}
                >
                  退室
                </Button>
              </Flex>
            </>
          ) : (
            <>
              <Button
                colorScheme="teal"
                variant="solid"
                size="lg"
                width={36}
                onClick={handleStatusChange}
              >
                入室
              </Button>
              <Flex ml="auto">
                <Button
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  width={36}
                  isDisabled={true}
                >
                  退室済
                </Button>
              </Flex>
            </>
          )}
        </Grid>
        
        )}
      </Grid>
      <TableContainer
        pb={14}
        pr={14}
        pl={14}
        mt={8}
        outlineOffset={2}
        overflowX="unset"
        overflowY="scroll"
        height="65vh"
      >
        <Table size="lg" border="2px" borderColor="gray.200" variant="simple">
          <Thead top={0}>
            <Tr bgColor="#E6EBED">
              <Th w="33%">出席者</Th>
              <Th w="33%">部屋</Th>
              <Th w="33%">入室時刻</Th>
            </Tr>
          </Thead>
          <Tbody outline="1px">
            {isFetching ? (
              <Td colSpan={5} textAlign="center">
                <Loading loadingItemText="出席者"></Loading>
              </Td>
            ) : attendeeList === null ? (
              <Tr>
                <Td colSpan={5} textAlign="center">
                  出席者はいません
                </Td>
              </Tr>
            ) : (
              attendeeList.map((attendee) => (
                <Tr key={attendee.user_id}>
                  <Td>
                    <Flex alignItems={"center"} gap={3}>
                      <Avatar
                        size={"md"}
                        src={attendee.avatar_img_path}
                        border="2px"
                        onClick={() =>
                          navigate("/profile", {
                            state: { userId: attendee.user_id },
                          })
                        }
                      ></Avatar>
                      {attendee.user_name}
                    </Flex>
                  </Td>
                  <Td>{attendee.place}</Td>
                  <Td>{decodeDate(attendee.entered_at)}</Td>
                </Tr>
              ))
            )}
          </Tbody>
        </Table>
      </TableContainer>
    </div>
  );
}

export default Home;
