import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import moment from 'moment';

import { EntCoolroom } from '../../api/models/EntCoolroom';
import { EntDeceasedReceive } from '../../api/models/EntDeceasedReceive';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [deceasedreceives, setDeceasedreceives] = useState<EntDeceasedReceive[]>([]);
 const [coolrooms, setCoolrooms] = useState<EntCoolroom[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getDeceasedreceives = async () => {
     const res = await api.listDeceasedreceive({ limit: 10, offset: 0 });
     setLoading(false);
     setDeceasedreceives(res);
     console.log(res);
   };
   getDeceasedreceives();

   const getCoolrooms = async () => {
    const res = await api.listCoolroom({ limit: 10, offset: 0 });
    setLoading(false);
    setCoolrooms(res);
  };
  getCoolrooms();

 }, [loading]);
 
 const deleteDeceasedreceives = async (id: number) => {
   const res = await api.deleteDeceasedreceive({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">ชื่อ-สกุล</TableCell>
           <TableCell align="center">ชื่อ-สกุล ของญาติ</TableCell>
           <TableCell align="center">หมายเลขห้องเก็บศพ</TableCell>
           <TableCell align="center">วัน/เดือน/ปี เวลา ที่เสียชีวิต</TableCell>
           <TableCell align="center">ประเภทของศพผู้เสียชีวิต</TableCell>
           <TableCell align="center">จัดการข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {deceasedreceives.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.patient.patientName}</TableCell>
             <TableCell align="center">{item.edges.relative.relativeName}</TableCell>
             <TableCell align="center">{item.edges.coolroom.coolroomName}</TableCell>
             <TableCell align="center">{moment(item.deathTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             {coolrooms.filter((filter:EntCoolroom) => filter.id === item.edges.coolroom.id).map((item2:any) => (
                  <TableCell align="center">{item2.edges.coolroomtype.coolroomtypeName}</TableCell>
              ))}
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteDeceasedreceives(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 ลบข้อมูล
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}