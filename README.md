## What is QLang
QLang is a xml language extension that allows anyone to create tests, quizes.

## Q&A
* Q) So QLang will help me to create tests what the difference between using google forms?
* A) QLang is a specification it is designed so any software can adopt QLang to define questions and answers.
________________
* Q) How easy is QLang and how much time does it take to learn it?
* A) It would take about maximum of 1 hour to learn QLang, Here's an example test written in QLang.
```xml
<?xml version="1.0" ?>
<qPack xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="./lang/def/qlang.xsd">
    <head text="Example test"/>
    <info author="Tarith"/>

    <qSet text="Fill in the blanks">
        <q text="le_rning">
            <choice correct="true" text="a"/>
            <choice correct="false" text="e"/>
            <choice correct="false" text="n"/>
        </q>
        <q text="bu__ness">
            <choice correct="false" text="is"/>
            <choice correct="true" text="si"/>
        </q>
    </qSet>

    <qSet text="Tyler had 30$ and he spent 5$ to buy some food then his father gave him 20$.">
        <q text="how much money dose he have now?">
            <choice correct="true" text="45$" />
            <choice correct="false" text="35$" />
            <choice correct="false" text="50$"/>
        </q>
    </qSet>
</qPack>
```