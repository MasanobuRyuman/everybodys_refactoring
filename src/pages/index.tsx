import type { NextPage } from "next";
import {
  Typography,
  Box,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  TableContainer,
  Button,
  Paper,
} from "@mui/material";

const tempData = ["question1", "question2"];

const Home: NextPage = () => {
  return (
    <div>
      <Box
        sx={{
          backgroundColor: "blue",
          height: 60,
        }}
      >
        <Typography
          sx={{
            color: "white",
            fontWeight: "bold",
            fontSize: 20,
            textAlign: "center",
            paddingTop: 2,
          }}
        >
          みんなのリファクタリング
        </Typography>
      </Box>
      <TableContainer
        component={Paper}
        sx={{ width: "80%", margin: "auto", marginTop: 5 }}
      >
        <Table>
          <TableHead>
            <TableRow>
              <TableCell
                sx={{
                  width: "80%",
                }}
              >
                問題名
              </TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tempData.map((value) => (
              <TableRow>
                <TableCell>{value}</TableCell>
                <TableCell>
                  <Button>移動</Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
};

export default Home;
