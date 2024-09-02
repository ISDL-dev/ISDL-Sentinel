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
import { inRoom } from "../../models/users/user";
import { attendeesListApi } from "../../api";
import { GetAttendeesList200ResponseInner } from "../../schema";
import { useEffect, useState } from "react";
import { useUser } from "../../userContext";
import { useNavigate } from "react-router-dom";

dayjs.locale("ja");

const decodeDate = (dateString: string) => {
  const date = dayjs(dateString);
  return `${dayjs(date).format("MM月DD日")}（${dayjs(date).format(
    "ddd"
  )}）${dayjs(date).format("HH時mm分")}`;
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

  useEffect(() => {
    fetchAttendeesList();
  }, []);

  return (
    <div>
      <Grid
        templateColumns="repeat(3, 1fr)"
        alignItems={"center"}
        w={"-moz-max-content"}
      >
        <h1 className="block mb-1 text-4xl font-bold text-gray-900 dark:text-white p-3 text-left">
          出席者一覧
        </h1>
        {authUser && (
          <Grid
            templateColumns="repeat(2, 1fr)"
            alignItems={"center"}
            w={"-moz-max-content"}
            column={3}
          >
            {authUser.status === inRoom ? (
              <>
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
                <Button
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  width={36}
                  isDisabled={true}
                >
                  退室済
                </Button>
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
                <Text pb={3} fontWeight={800} fontSize={20}>
                  出席者を読込中です
                </Text>
                <Spinner
                  thickness="4px"
                  speed="0.65s"
                  emptyColor="gray.200"
                  color="teal.500"
                  size="xl"
                  alignContent={"center"}
                />
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
