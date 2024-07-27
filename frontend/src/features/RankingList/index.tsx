import {
  Avatar,
  AvatarBadge,
  Box,
  Button,
  Card,
  CardBody,
  Center,
  Flex,
  Grid,
  GridItem,
  Heading,
  Image,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { Top3Icon } from "../RankIcon/Top3Icon";
import { Top10Icon } from "../RankIcon/Top10Icon";

type RankingList = {
  user_id: number;
  user_name: string;
  attendance_days: number;
  stay_time: string;
  grade: number;
  avatar_id: number;
  avatar_img_path: string;
};
const rankingList: RankingList[] = [];
for (var i = 1; i < 11; i++) {
  rankingList.push({
    user_id: i,
    user_name: `user${i}`,
    attendance_days: i * 4,
    grade: i,
    avatar_id: 1,
    stay_time: "09:00:00",
    avatar_img_path: "default1.png",
  });
}
export const RankingList = (placeholder: { placeholder: string }) => {
  return (
    <>
      <Grid
        templateColumns="repeat(3, 1fr)"
        templateRows="1fr 2fr 1fr 1fr 10fr"
        justifyItems={"center"}
        placeItems={"center"}
        alignItems={"center"}
        height={"20%"}
        w={"-moz-max-content"}
        column={3}
        row={4}
      >
        <Top3Icon rank="1" color="#ffd700"></Top3Icon>
        <Top3Icon rank="2" color="#c9caca"></Top3Icon>
        <Top3Icon rank="3" color="#b87333"></Top3Icon>
        <Image
          src={`./avatar/default1.png`}
          alt={"default"}
          boxSize="64px"
          cursor="pointer"
        />
        <Image
          src={`./avatar/default1.png`}
          alt={"default"}
          boxSize="64px"
          cursor="pointer"
        />
        <Image
          src={`./avatar/default1.png`}
          alt={"default"}
          boxSize="64px"
          cursor="pointer"
        />
        <Text>ユーザ1</Text>
        <Text>ユーザ2</Text>
        <Text>ユーザ3</Text>
        <Heading fontSize="2xl" color="#fa8072">
          240時間
        </Heading>
        <Heading fontSize="2xl" color="#fa8072">
          210時間
        </Heading>
        <Heading fontSize="2xl" color="#fa8072">
          160時間
        </Heading>
        <GridItem rowSpan={1} colSpan={3}>
          <TableContainer
            pb={14}
            pr={14}
            pl={14}
            pt={3}
            overflowX="auto"
            overflowY="scroll"
            height="45vh"
            width="120vh"
          >
            <Table size="sm">
              <Tbody>
                {(() => {
                  const rankUserRender = [];
                  for (var i = 4; i < rankingList.length; i++) {
                    rankUserRender.push(
                      <Tr key={rankingList[i].user_id}>
                        <Td textAlign="center" w={15}>
                          <Card>
                            <CardBody>
                              <Grid
                                templateColumns="1fr 1fr 1fr 5fr"
                                alignItems={"center"}
                                height={"20%"}
                                w={"-moz-max-content"}
                                column={3}
                                row={4}
                              >
                                <Top10Icon rank={`${i}`}></Top10Icon>
                                <Image
                                  src={`./avatar/${rankingList[i].avatar_img_path}`}
                                  alt={`${rankingList[i].avatar_id}`}
                                  boxSize="50px"
                                  cursor="pointer"
                                />
                                <Text fontSize="large">
                                  {rankingList[i].user_name}
                                </Text>
                                <GridItem colEnd={6}>
                                  <Heading fontSize="2xl" color="#fa8072">
                                    {rankingList[i].stay_time}
                                  </Heading>
                                </GridItem>
                              </Grid>
                            </CardBody>
                          </Card>
                        </Td>
                      </Tr>
                    );
                  }
                  return rankUserRender;
                })()}
              </Tbody>
            </Table>
          </TableContainer>
        </GridItem>
      </Grid>
    </>
  );
};
