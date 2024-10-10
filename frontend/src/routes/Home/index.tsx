import {
  Avatar,
  Button,
  Flex,
  Grid,
  Box,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  Icon,
} from "@chakra-ui/react";
import { FaBed } from "react-icons/fa";
import "./Home.css";
import dayjs from "dayjs";
import "dayjs/locale/ja";
import { inRoom, outRoom, overnight } from "../../models/users/user";
import { attendeesListApi, profileApi } from "../../api";
import { GetAttendeesList200ResponseInner } from "../../schema";
import { useEffect, useState } from "react";
import { useUser } from "../../userContext";
import { useNavigate } from "react-router-dom";
import { Loading } from "../../features/Loading/Loading";

const buttonWidth = {
  base: "100px",
  md: "150px",
};

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
  return hour >= 20 && hour < 24;
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

      if (authUser) {
        const user = await profileApi.getUserById(authUser.user_id);
        setAuthUser({
          ...authUser,
          status: user.data.status,
        });
      }

      const response = await attendeesListApi.getAttendeesList();
      setAttendeeList(response.data);
      setIsFetching(false);
    } catch (error) {
      console.log(error);
    }
  };

  const handleStatusChange = async (requestedStatus: string) => {
    if (!authUser) return;
    try {
      const user = await attendeesListApi.putStatus({
        user_id: authUser.user_id,
        status: requestedStatus,
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
        templateColumns="repeat(1, 1fr)"
        w="100%"
        mt={{ base: 20, md: 0 }}
        p={6}
      >
        <Flex justifyContent="center" alignItems="center" w="100%">
          <Box flex="1">
            <Text
              fontSize={{ base: "2xl", md: "4xl" }}
              fontWeight="bold"
              color="gray.900"
              mb={3}
              whiteSpace="nowrap"
            >
              出席者一覧
            </Text>
          </Box>
        </Flex>

        {authUser && (
          <Flex
            justifyContent={{ base: "center", md: "flex-end" }}
            alignItems="center"
            gap={4}
            mt={4}
            flexWrap="nowrap"
          >
            {authUser.status === overnight ? (
              <>
                <Button
                  w={buttonWidth}
                  colorScheme="cyan"
                  variant="solid"
                  size="lg"
                  isDisabled
                >
                  宿泊済
                </Button>
                <Button
                  w={buttonWidth}
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  isDisabled
                >
                  入室済
                </Button>
                <Button
                  w={buttonWidth}
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  onClick={() => handleStatusChange(outRoom)}
                >
                  退室
                </Button>
              </>
            ) : authUser.status === inRoom ? (
              <>
                {isBetween8PMandMidnight() && (
                  <Button
                    w={buttonWidth}
                    colorScheme="cyan"
                    variant="solid"
                    size="lg"
                    onClick={() => handleStatusChange(overnight)}
                  >
                    宿泊
                  </Button>
                )}
                <Button
                  w={buttonWidth}
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  isDisabled
                >
                  入室済
                </Button>
                <Button
                  w={buttonWidth}
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  onClick={() => handleStatusChange(outRoom)}
                >
                  退室
                </Button>
              </>
            ) : (
              <>
                <Button
                  w={buttonWidth}
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  onClick={() => handleStatusChange(inRoom)}
                >
                  入室
                </Button>
                <Button
                  w={buttonWidth}
                  colorScheme="teal"
                  variant="solid"
                  size="lg"
                  isDisabled
                >
                  退室済
                </Button>
              </>
            )}
          </Flex>
        )}
      </Grid>

      <TableContainer
        pb={14}
        pr={{ base: 4, md: 14 }}
        pl={{ base: 4, md: 14 }}
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
                      <Box position="relative">
                        <Avatar
                          size={"md"}
                          src={attendee.avatar_img_path}
                          border="2px"
                          onClick={() =>
                            navigate("/profile", {
                              state: { userId: attendee.user_id },
                            })
                          }
                        />
                        {attendee.status === overnight && (
                          <Icon
                            as={FaBed}
                            color="cyan.500"
                            position="absolute"
                            top={-1}
                            right={-1}
                          />
                        )}
                      </Box>
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
