import React, { useState, useEffect } from 'react';
import Axios from 'axios';
import {
    Button,
    Form,
    Container,
    Modal,
    ModalHeader,
    ModalTitle,
    ModalBody,
    FormGroup,
    FormLabel, FormControl
} from 'react-bootstrap';
import Entry from './single-entry.components'


const Entries = () => {
    const [entries, setEntries ] = useState([])
    const [refreshData, setRefreshData ] = useState(false)
    const [changeEntry, setChangeEntry] = useState({"change":false, "id":0})
    const [addNewEntry, setAddNewEntry] = useState(false)
    const [newEntry, setNewEntry] = useState({"dish":"", "ingredients":"", "Calories":"", "Protein":"", "Fats":0})

    useEffect(()=>{
        getAllEntries();
    }, [])

    if(refreshData){
        setRefreshData(false);
        getAllEntries();
    }

    return (
        <div>
            <Container>
                <Button onClick={()=> setAddNewEntry(true)}>Track Today's Calories</Button>
            </Container>
            <Container>
                { entries !== null && entries.map((entry, i) => (
                    <Entry entryData={entry} deleteSingleEntry={deleteSingleEntry} setChangeEntry={setChangeEntry}/>
                ))}
            </Container>

            <Modal show={addNewEntry} onHide={()=>setAddNewEntry(false)} centered={}>
                <ModalHeader closeButton>
                    <ModalTitle>Add Calorie Entry</ModalTitle>
                </ModalHeader>
                <ModalBody>
                    <FormGroup>
                        <FormLabel>dish</FormLabel>
                        <FormControl onChange={(event)=>{newEntry.dish = event.target.value }}></FormControl>

                        <FormLabel>ingredients</FormLabel>
                        <FormControl onChange={(event)=>{newEntry.ingredients = event.target.value }}></FormControl>

                        <FormLabel>calories</FormLabel>
                        <FormControl onChange={(event)=>{newEntry.calories = event.target.value }}></FormControl>

                        <FormLabel>protein</FormLabel>
                        <FormControl onChange={(event)=>{newEntry.protein = event.target.value }}></FormControl>

                        <FormLabel>fats</FormLabel>
                        <FormControl onChange={(event)=>{newEntry.fats = event.target.value }}></FormControl>
                    </FormGroup>
                    <Button onClick={()=>addSingleEntry()}>Add</Button>
                    <Button onClick={()=>setAddNewEntry(false)}>Cancel</Button>
                </ModalBody>
            </Modal>

        </div>
    );

    function getAllEntries(){
        var url = "http://localhost:8001/entries"
        axios.get(url, {
            responseType:'json'
        }).then(resonse => {
            if(response.status === 200){
                setRefreshData(true)
            }
        })
    }

    function addSingleEntry(){
        setAddNewEntry(false){
            var url = "http://localhost:8001/entry/create"
            axios.post(url, {
                "ingredients": newEntry.ingredients,
                "dish": newEntry.dish,
                "calories": newEntry.calories,
                "protein": newEntry.protein,
                "fats": parseFloat(newEntry.fats)
            }).then(response=> {
                if(response.status === 200){
                    setRefreshData(true)
                }
            })
        }
    }

    function deleteSingleEntry(id){
        var url = "http://localhost:8001/entry/delete" + id
        axios.delete(url, {

        }).then(response=> {
            if(response.status === 200){
                setRefreshData(true)
            }
        })
    }
}





