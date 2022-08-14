import React, { useState, useEffect } from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import {Button, Card, Row, Col} from 'react-bootstrap';

const Entry = ({entryData, deleteSingleEntry, setChangeEntry}) => {
    return (
        <Card>
            <Row>
                <Col>Dish: {entryData !== undefined && entryData.dish}</Col>
                <Col>Ingredients: {entryData !== undefined && entryData.ingredients}</Col>
                <Col>Calories: {entryData !== undefined && entryData.calories}</Col>
                <Col>Protein: {entryData !== undefined && entryData.protein}</Col>
                <Col>Fats: {entryData !== undefined && entryData.fats}</Col>
                <Col><Button onClick={()=> deleteSingleEntry(entryData._id)}>Update Entry</Button></Col>
                <Col><Button onClick={()=> changeEntry()}>Delete Entry</Button></Col>
            </Row>
        </Card>
    )

    function changeEntry(){
        setChangeEntry({
            "change": true,
            "id": entryData._id
        })
    }

}

export default Entry
