import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  HttpCode,
  UseGuards,
} from '@nestjs/common';
import { EventsService } from '@app/core/events/events.service';
import { CreateEventRequest } from './request/create-event.request';
import { UpdateEventRequest } from './request/update-event.request';
import { ReserveSpotRequest } from './request/reserve-spot.request';
import { AuthGuard } from '@app/core/auth/auth.guard';
<<<<<<< HEAD
=======
import { ReserveSpotResponse } from './response/reserve-spot.response';
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)

@Controller('events')
export class EventsController {
  constructor(private readonly eventsService: EventsService) {}

  @Post()
  create(@Body() createEventRequest: CreateEventRequest) {
    return this.eventsService.create(createEventRequest);
  }

  @Get()
  findAll() {
    return this.eventsService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.eventsService.findOne(id);
  }

  @Patch(':id')
  update(
    @Param('id') id: string,
    @Body() updateEventRequest: UpdateEventRequest,
  ) {
    return this.eventsService.update(id, updateEventRequest);
  }

  @HttpCode(204)
  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.eventsService.remove(id);
  }

  @UseGuards(AuthGuard)
  @Post(':id/reserve')
<<<<<<< HEAD
  reserveSpots(
    @Body() reserveRequest: ReserveSpotRequest,
    @Param('id') eventId: string,
  ) {
    return this.eventsService.reserveSpot({ ...reserveRequest, eventId });
=======
  async reserveSpots(
    @Body() reserveRequest: ReserveSpotRequest,
    @Param('id') eventId: string,
  ) {
    const tickets = await this.eventsService.reserveSpot({
      ...reserveRequest,
      eventId,
    });
    return new ReserveSpotResponse(tickets);
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
  }
}
