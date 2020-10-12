import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { EntCoolroom } from '../../api/models/EntCoolroom';
import { EntCoolroomType } from '../../api/models/EntCoolroomType';
import { EntRelative } from '../../api/models/EntRelative';
import { EntPatient } from '../../api/models/EntPatient';
import { EntDeceasedReceive } from '../../api/models/EntDeceasedReceive';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const initialUserState = {
  dOCTORNAME: 'นายแพทย์กำลัง นอนพอดี',
  dOCTOREMAIL: 'Gumlung@gmail.com',
};

const username = { givenuser: 'นายแพทย์กำลัง นอนพอดี' };

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบผู้เสียชีวิต' };
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [coolrooms, setCoolrooms] = useState<EntCoolroom[]>([]);
  const [coolroomtypes, setCoolroomTypes] = useState<EntCoolroomType[]>([]);
  const [relatives, setRelatives] = useState<EntRelative[]>([]);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [deceasedreceives, setDeceasedreceives] = useState<EntDeceasedReceive[]>([]);

  const [loading, setLoading] = useState(true);
  const [deathtime, setDeathtime] = useState(String);

  const [coolroomid, setcoolroom] = useState(Number);
  const [typeid, setDeadStatus] = useState(Number);
  const [relativeid, setRelative] = useState(Number);
  const [patientid, setPatient] = useState(Number);

  useEffect(() => {
    const getCoolroomTypes = async () => {
      const res = await api.listCoolroomtype({ limit: 10, offset: 0 });
      setLoading(false);
      setCoolroomTypes(res);
    };
    getCoolroomTypes();

    const getCoolrooms = async () => {
      const res = await api.listCoolroom({ limit: 10, offset: 0 });
      setLoading(false);
      setCoolrooms(res);
      console.log(res);
    };
    getCoolrooms();

    const getRelatives = async () => {
      const res = await api.listRelative({ limit: 10, offset: 0 });
      setLoading(false);
      setRelatives(res);
    };
    getRelatives();

    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(res);
      console.log(res);
    };
    getPatients();

  }, [loading]);

  const getDeceasedReceives = async () => {
    const res = await api.listDeceasedreceive({ limit: 10, offset: 0 });
    setDeceasedreceives(res);
  };

  const handleDeathtimeChange = (event: any) => {
    setDeathtime(event.target.value as string);
  };

  const CoolroomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setcoolroom(event.target.value as number);
  };

  const DeathtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDeadStatus(event.target.value as number);
  };

  const RelativehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRelative(event.target.value as number);
  };

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatient(event.target.value as number);
  };

  const CreateDeceasedReceive = async () => {
    const deceasedreceive = {
      patientID: patientid,
      coolroomID: coolroomid,
      relativeID: relativeid,
      userID: 1,
      deathtime: deathtime + ":00+07:00"

    };
    console.log(deceasedreceive);
    const res: any = await api.createDeceasedreceive({ deceasedreceive: deceasedreceive });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลผู้เสียชีวิต">
          <Typography align="left" style={{ marginRight: 16, color: "#00eeff" }}>
            {username.givenuser}
          </Typography>
          <div>
            <Button variant="contained" color="primary">
              ออกจากระบบ
         </Button>
          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="patient-label">ชื่อ-สกุล ผู้เสียชีวิต</InputLabel>
                <Select
                  labelId="patient-label"
                  id="patient"
                  value={patientid}
                  onChange={PatienthandleChange}
                  style={{ width: 400 }}
                >
                {patients.map((item: EntPatient) => (
                  <MenuItem value={item.id}>{item.patientName}</MenuItem>
                ))}
                </Select>
              </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="relative-label">ชื่อ-สกุล ญาติผู้เสียชีวิต</InputLabel>
                <Select
                  labelId="relative-label"
                  id="relative"
                  value={relativeid}
                  onChange={RelativehandleChange}
                  style={{ width: 400 }}
                >
                {relatives.map((item: any) => (
                  <MenuItem value={item.id}>{item.relativeName}</MenuItem>
                ))}
                </Select>
              </FormControl>
            </div>

            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="status-death">ประเภทศพผู้เสียชีวิต</InputLabel>
              <Select
                labelId="status-death"
                id="statusdeath"
                value={typeid}
                onChange={DeathtypehandleChange}
                style={{ width: 200 }}
              >
                {coolroomtypes.map((item: EntCoolroomType) => (
                  <MenuItem value={item.id}>{item.coolroomtypeName}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="coolroom-label">หมายเลขห้องเย็น</InputLabel>
                <Select
                  labelId="coolroom-label"
                  id="coolroom"
                  value={coolroomid}
                  onChange={CoolroomhandleChange}
                  style={{ width: 200 }}
                >
                  {coolrooms.filter((setfilterid:any) => setfilterid.edges.coolroomtype.id === typeid).map(item => (
                  <MenuItem value={item.id}>{item.coolroomName}</MenuItem>                 
                ))}
                </Select>
              </FormControl>
            
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="deathtime"
                  label="ว/ด/ป เวลาเสียชีวิต"
                  type="datetime-local"
                  value={deathtime}
                  onChange={handleDeathtimeChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
            </div>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateDeceasedReceive();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}