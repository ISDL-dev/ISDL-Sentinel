import {
  Card,
  CardBody,
  Grid,
  GridItem,
  Heading,
  Image,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
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
        templateRows="1fr 2fr 1fr 1fr 12fr"
        justifyItems={"center"}
        placeItems={"center"}
        alignItems={"center"}
        height={"20%"}
        w={"-moz-max-content"}
        column={3}
        row={4}
      >
        {(() => {
          const rankUserRender = [];
          for (var i = 0; i < 3; i++) {
            rankUserRender.push(
              <GridItem rowSpan={4} colSpan={1} textAlign={"center"}>
                <Top3Icon rank={i}></Top3Icon>
                <Image
                  src={`./avatar/default1.png`}
                  alt={"default"}
                  boxSize={{
                    base: "48px",
                    md: "64px",
                  }}
                  cursor="pointer"
                  ml={5}
                />
                <Text>{rankingList[i].user_name}</Text>
                <Heading fontSize={{ base: "xl", md: "2xl" }} color="#fa8072">
                  {rankingList[i].stay_time}
                </Heading>
              </GridItem>
            );
          }
          return rankUserRender;
        })()}

        <GridItem rowSpan={1} colSpan={3}>
          <TableContainer
            pb={14}
            pr={14}
            pl={14}
            pt={3}
            overflowX="auto"
            overflowY="scroll"
            height={{
              base: "50vh",
              md: "48vh",
            }}
            width={{
              base: "120vw",
              md: "60vw",
            }}
          >
            <Table size="sm">
              <Tbody>
                {(() => {
                  const rankUserRender = [];
                  for (var i = 3; i < rankingList.length; i++) {
                    rankUserRender.push(
                      <Tr key={rankingList[i].user_id}>
                        <Td textAlign="center" w={15}>
                          <Card>
                            <CardBody>
                              <Grid
                                templateColumns="auto auto 1fr auto"
                                alignItems={"center"}
                                height={"20%"}
                                w={"-moz-max-content"}
                                gap={4}
                              >
                                <GridItem
                                  justifySelf="start"
                                  marginRight={{ base: 12, md: 20 }}
                                  paddingTop={14}
                                >
                                  <Top10Icon rank={`${i + 1}`} />
                                </GridItem>
                                <GridItem
                                  justifySelf="start"
                                  marginRight={{ base: 0, md: 8 }}
                                >
                                  <Image
                                    src={`./avatar/${rankingList[i].avatar_img_path}`}
                                    alt={`${rankingList[i].avatar_id}`}
                                    boxSize={{
                                      base: "36px",
                                      md: "50px",
                                    }}
                                    cursor="pointer"
                                  />
                                </GridItem>
                                <GridItem justifySelf="start">
                                  <Text fontSize="large">
                                    {rankingList[i].user_name}
                                  </Text>
                                </GridItem>
                                <GridItem justifySelf="end">
                                  <Heading
                                    fontSize={{ base: "xl", md: "2xl" }}
                                    color="#fa8072"
                                  >
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
