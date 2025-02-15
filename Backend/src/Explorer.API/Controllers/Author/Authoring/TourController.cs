﻿using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.MicroserviceDtos;
using Explorer.Tours.API.Public.Authoring;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Text.Json;
using Newtonsoft.Json;
using System.Threading.Tasks;
using System.Numerics;
using Explorer.Tours.Core.Domain.Tours;
using Newtonsoft.Json.Serialization;
using System.Text;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace Explorer.API.Controllers.Author.Authoring
{
    [Authorize(Policy = "authorPolicy")]
    [Route("api/tourManagement/tour")]
    public class TourController : BaseApiController
    {
        private readonly ITourService _tourService;
        private readonly IHttpClientFactory _factory;

        public TourController(ITourService tourService, IHttpClientFactory factory)
        {
            _tourService = tourService;
            _factory = factory;
        }

        [HttpGet]
        public async Task<PagedResult<TourDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            //var result = _tourService.GetPaged(page, pageSize);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.GetAsync("tours");
            var jsonResponse = await response.Content.ReadAsStringAsync();
            List<TourDto> tourDtos = System.Text.Json.JsonSerializer.Deserialize<List<TourDto>>(jsonResponse);
            int totalCount = GetTotalCountFromHeaders(response);

            var pagedResult = new PagedResult<TourDto>(tourDtos, totalCount);

            return pagedResult;
        }

        [HttpPost]
        public async Task<ActionResult<TourDto>> Create([FromBody] TourDto tourDto)
        {
            //var result = _tourService.Create(tourDto);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PostAsJsonAsync("/tours/create", tourDto);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();
            var createdTour = System.Text.Json.JsonSerializer.Deserialize<TourDto>(jsonResponse);

            return Ok(createdTour);
        }

        [HttpPut("{id:int}")]
        public async Task<ActionResult<TourDto>> Update(int id, [FromBody] TourDto tourDto)
        {
            //var result = _tourService.Update(tourDto);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("/tours/update", tourDto);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }
            var jsonResponse = await response.Content.ReadAsStringAsync();
            var updatedTour = System.Text.Json.JsonSerializer.Deserialize<TourDto>(jsonResponse);

            return Ok(updatedTour);
        }


        [AllowAnonymous]
        [HttpGet("{id:int}")]
        public async Task<ActionResult<TourDto>> Get(int id)
        {
            //var result = _tourService.Get(id);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.GetAsync("tours/" + id);
            var jsonResponse = await response.Content.ReadAsStringAsync();
            TourDto tourDto = System.Text.Json.JsonSerializer.Deserialize<TourDto>(jsonResponse);

            return Ok(tourDto);
        }

        [HttpPut("publish/{id:int}")]
        public async Task<ActionResult<TourDto>> Publish(int id, [FromBody] int authorId)
        {
            //var result = _tourService.Publish(id, authorId);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("/tours/publish/" + id, authorId);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }
            var jsonResponse = await response.Content.ReadAsStringAsync();
            var publishedTour = System.Text.Json.JsonSerializer.Deserialize<TourDto>(jsonResponse);

            return Ok(publishedTour);
        }

        [HttpPut("archive/{id:int}")]
        public async Task<ActionResult<TourDto>> Archive(int id, [FromBody] int authorId)
        {
            //var result = _tourService.Archive(id, authorId);
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("/tours/archive/" + id, authorId);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }
            var jsonResponse = await response.Content.ReadAsStringAsync();
            var archivedTour = System.Text.Json.JsonSerializer.Deserialize<TourDto>(jsonResponse);

            return Ok(archivedTour);
        }

        [HttpGet("author")]
        public async Task<PagedResult<TourDto>> GetAllByAuthorId([FromQuery] int authorId, [FromQuery] int page, [FromQuery] int pageSize)
        {
            var client = _factory.CreateClient("toursMicroservice");
            using HttpResponseMessage response = await client.GetAsync($"tours/author/" + authorId);
            var jsonResponse = await response.Content.ReadAsStringAsync();
            var tourDtos = System.Text.Json.JsonSerializer.Deserialize<List<TourDto>>(jsonResponse);
            int totalCount = GetTotalCountFromHeaders(response);

            var pagedResult = new PagedResult<TourDto>(tourDtos,totalCount);

            return pagedResult;
        }

        private int GetTotalCountFromHeaders(HttpResponseMessage response)
        {
            if (response.Headers.TryGetValues("X-Total-Count", out var values))
            {
                int.TryParse(values.FirstOrDefault(), out int totalCount);
                return totalCount;
            }
            return 0;
        }


    }
}
